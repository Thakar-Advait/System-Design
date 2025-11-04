package ui

type DarkButton struct{}
type DarkTextbox struct{}

func (DarkButton) Render() string {
	return "Dark Button"
}

func (DarkTextbox) Render() string {
	return "Dark Textbox"
}

type DarkFactory struct{}

func (DarkFactory) CreateButton() Button {
	return DarkButton{}
}

func (DarkFactory) CreateTextbox() Textbox {
	return DarkTextbox{}
}
