package ui

import (
	"async_communication/internal/model"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	listView sessionState = iota
	inputView
)

type SurveyModel struct {
	listIsReady bool
	state       sessionState
	input       textinput.Model
	list        list.Model
	locations   *model.Locations
	selectedHit model.Hit
}

func (s SurveyModel) Init() tea.Cmd {
	return textinput.Blink
}

func (s SurveyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmdList []tea.Cmd
	var inputCmd tea.Cmd
	var listCmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if s.input.Focused() {
				cmdList = append(cmdList, RequestLocations(s.input.Value()))
				s.input.Blur()
			} else {
				s.selectedHit = s.locations.Hits[s.list.Cursor()]
				cmdList = append(cmdList, CmdSwitchToDetails(s.selectedHit))
			}
		case tea.KeyEscape:
			s.input.Focus()
		}
	case LocationsLoaded:
		s.listIsReady = true
		s.locations = msg.Locations
		createList(&s.list, msg.Locations)
	}

	s.input, inputCmd = s.input.Update(msg)
	s.list, listCmd = s.list.Update(msg)
	cmdList = append(cmdList, inputCmd)
	cmdList = append(cmdList, listCmd)

	return s, tea.Batch(cmdList...)
}

func (s SurveyModel) View() string {
	if !s.listIsReady {
		return s.input.View()
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, s.input.View(), s.list.View())
}

func InitSurvey() SurveyModel {
	t := textinput.New()
	t.Placeholder = "What location is on your mind today?"
	t.Focus()
	t.CharLimit = 156
	t.Width = 40

	l := list.New(nil, list.NewDefaultDelegate(), 60, 20)
	l.DisableQuitKeybindings()
	l.SetShowHelp(false)

	return SurveyModel{
		listIsReady: false,
		state:       inputView,
		input:       t,
		list:        l,
		locations:   nil,
	}
}

func createList(l *list.Model, locations *model.Locations) {
	var items []list.Item

	for _, value := range locations.Hits {
		name := value.Name

		if len(name) == 0 {
			name = "Без имени"
		}

		items = append(items, ListItem{
			title:       fmt.Sprintf("%s - %s", name, value.City),
			description: value.Country,
		})
	}

	l.SetItems(items)
	l.Title = "Found locations"
}

func CmdSwitchToDetails(selectedHit model.Hit) tea.Cmd {
	return func() tea.Msg {
		return SwitchToDetails{Hit: selectedHit}
	}
}
