package gameengine

import "fmt"

type PlatformSvc interface {
	CreateButton() Button
	CreateTextbox() Textbox
	CreateDialog() Dialog
}

type Factory interface {
	NewEngineBuilder() Builder
	Service() PlatformSvc
}

func NewPlatformFactory(vertical string) (Factory, error) {
	switch vertical {
	case "pc":
		return &PCFactory{}, nil
	case "mobile":
		return &MobileFactory{}, nil
	case "console":
		return &ConsoleFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported platform type")
	}
}
