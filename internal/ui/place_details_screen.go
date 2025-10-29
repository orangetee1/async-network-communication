package ui

import (
	"async_communication/internal/model"
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PlaceDetailsModel struct {
	listLoaded   bool
	weather      string
	selectedInfo string
	list         list.Model
	descriptions map[string]string
}

func (p PlaceDetailsModel) Init() tea.Cmd {
	return nil
}

func (p PlaceDetailsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var listCmd tea.Cmd

	switch msg := msg.(type) {
	case SwitchToDetails:
		cmds = append(cmds, RequestWeather(msg.Hit.Latitude, msg.Hit.Longitude))
		cmds = append(cmds, RequestPlaces(msg.Hit.Latitude, msg.Hit.Longitude))
	case WeatherLoaded:
		p.weather = createWeather(msg.Weather)
	case PlacesLoaded:
		p.listLoaded = true
		createListPlaces(&p.list, msg.Places)
		cmds = append(cmds, createInfoRequests(msg.Places))
	case PlaceInfoLoaded:
		p.descriptions[msg.Info.Features[0].Properties.Id] = createDescription(msg.Info)
	}

	p.list, listCmd = p.list.Update(msg)

	if p.listLoaded {
		val, ok := p.descriptions[p.list.SelectedItem().(ListItem).Id()]
		if ok {
			p.selectedInfo = val
		} else {
			p.selectedInfo = "Not loaded yet"
		}
	}

	cmds = append(cmds, listCmd)

	return p, tea.Batch(cmds...)
}

func (p PlaceDetailsModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top, p.weather, p.list.View(),
		lipgloss.JoinVertical(lipgloss.Top, PlaceInfoHeader.Render(), p.selectedInfo))
}

func InitPlaceDetails() PlaceDetailsModel {
	l := list.New(nil, list.NewDefaultDelegate(), 40, 20)
	l.DisableQuitKeybindings()
	l.SetShowHelp(false)

	return PlaceDetailsModel{
		list:         l,
		selectedInfo: "Nothing selected",
		descriptions: make(map[string]string),
	}
}

func createListPlaces(l *list.Model, places *model.Places) {
	var items []list.Item

	for _, value := range places.Features {
		items = append(items, ListItem{
			id:          value.Properties.Id,
			title:       value.Properties.Name,
			description: value.Properties.Country,
		})
	}

	l.SetItems(items)
	l.Title = "Interesting places"
}

func createWeather(weather *model.Weather) string {
	var desc string

	for _, value := range weather.Description {
		desc += fmt.Sprintf("%s %s", value.Description, value.Main)
	}

	return lipgloss.JoinVertical(lipgloss.Top, WeatherHeader.Render(),
		fmt.Sprintf("Description: %s\nHumidity: %d\nWind speed: %f\n",
			desc, weather.Values.Humidity, weather.Wind.Speed)) // todo: fix, add metrics, make it table
}

func createInfoRequests(places *model.Places) tea.Cmd {
	var cmds []tea.Cmd

	for _, val := range places.Features {
		cmds = append(cmds, RequestPlaceInfo(val.Properties.Id))
	}

	return tea.Batch(cmds...)
}

func createDescription(placeInfo *model.PlaceInfo) string {
	props := placeInfo.Features[0].Properties

	return fmt.Sprintf("Name: %s\nCity: %s\nStreet: %s\nHousenumber: %s\nPhone: %s",
		props.Name, props.City, props.Street, props.HouseNumber, props.Contact.Phone)
}
