package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"os"
)

type sessionState int

const (
	surveyView sessionState = iota
	detailsView
)

type MainModel struct {
	state sessionState

	surveyModel       SurveyModel
	placeDetailsModel PlaceDetailsModel

	currentError string

	logger *os.File
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(m.surveyModel.Init(), m.placeDetailsModel.Init())
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdList []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case ErrorChanged:
		m.currentError = msg.Error

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}

	case SwitchToDetails:
		m.state = detailsView
	case SwitchToSurvey:
		m.state = surveyView
	}

	switch m.state {
	case surveyView:
		newSurvey, surveyCmd := m.surveyModel.Update(msg)
		survey, ok := newSurvey.(SurveyModel)
		if !ok {
			log.Fatal()
		}
		m.surveyModel = survey
		cmd = surveyCmd
	case detailsView:
		newDetails, detailsCmd := m.placeDetailsModel.Update(msg)
		details, ok := newDetails.(PlaceDetailsModel)
		if !ok {
			log.Fatal()
		}
		m.placeDetailsModel = details
		cmd = detailsCmd
	}

	cmdList = append(cmdList, cmd)

	return m, tea.Batch(cmdList...)
}

func (m MainModel) View() string {
	var view string

	switch m.state {
	case surveyView:
		view = m.surveyModel.View()
	case detailsView:
		view = m.placeDetailsModel.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top, view, ErrorStyle.Render(m.currentError))
}

func InitMain(logger *os.File) MainModel {
	log.SetOutput(logger)

	return MainModel{
		surveyModel:       InitSurvey(),
		placeDetailsModel: InitPlaceDetails(),
		currentError:      "No error",
		state:             surveyView,
		logger:            logger,
	}
}
