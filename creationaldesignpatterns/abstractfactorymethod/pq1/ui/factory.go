package ui

import "fmt"

type UIFactory interface {
	CreateButton() Button
	CreateTextbox() Textbox
}

func NewUIFactory(vertical string) (UIFactory, error) {
	switch vertical {
	case "light":
		return LightFactory{}, nil
		// panic("to be implemented")
	case "dark":
		return DarkFactory{}, nil
		//panic("to be implemented")
	default:
		return nil, fmt.Errorf("unsupported UI vertical")
	}
}
