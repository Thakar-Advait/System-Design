package ui

type LightButton struct{}
type LightTextbox struct{}

func (LightButton) Render() string {
	return "Light Button"
}

func (LightTextbox) Render() string {
	return "Light Textbox"
}

type LightFactory struct{}

func (LightFactory) CreateButton() Button {
	return LightButton{}
}

func (LightFactory) CreateTextbox() Textbox {
	return LightTextbox{}
}
