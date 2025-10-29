package ui

type ListItem struct {
	id          string
	title       string
	description string
}

func (l ListItem) Id() string {
	return l.id
}

func (l ListItem) FilterValue() string {
	return l.title
}

func (l ListItem) Title() string {
	return l.title
}

func (l ListItem) Description() string {
	return l.description
}
