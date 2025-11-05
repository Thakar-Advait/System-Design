package gameengine

import "fmt"

// ConsoleBuilder implements Builder for Console platform
type ConsoleBuilder struct {
	Shadows         bool
	ParticleEffects bool
	Sound           bool
	Platform        string
}

func (console *ConsoleBuilder) EnableShadows() Builder {
	console.Shadows = true
	return console
}

func (console *ConsoleBuilder) EnableParticleEffects() Builder {
	console.ParticleEffects = true
	return console
}

func (console *ConsoleBuilder) EnableSound() Builder {
	console.Sound = true
	return console
}

func (console *ConsoleBuilder) Build() *PhysicsEngine {
	engine := GetInstance()
	engine.Shadows = console.Shadows
	engine.ParticleEffects = console.ParticleEffects
	engine.Sound = console.Sound
	engine.Platform = console.Platform
	return engine
}

// ConsoleFactory implements Factory for Console platform
type ConsoleFactory struct{}

func (*ConsoleFactory) NewEngineBuilder() Builder {
	return &ConsoleBuilder{
		Platform: "console",
	}
}

func (*ConsoleFactory) Service() PlatformSvc {
	return &ConsoleSvc{}
}

// ConsoleSvc implements PlatformSvc for Console platform UI components
type ConsoleSvc struct{}

func (svc *ConsoleSvc) CreateButton() Button {
	return &ConsoleButton{}
}

func (svc *ConsoleSvc) CreateTextbox() Textbox {
	return &ConsoleTextbox{}
}

func (svc *ConsoleSvc) CreateDialog() Dialog {
	return &ConsoleDialog{}
}

// Console UI Component implementations
type ConsoleButton struct{}

func (*ConsoleButton) Render() {
	fmt.Println("Rendered a Console Button")
}

type ConsoleTextbox struct{}

func (*ConsoleTextbox) Render() {
	fmt.Println("Rendered a Console Textbox")
}

type ConsoleDialog struct{}

func (*ConsoleDialog) Render() {
	fmt.Println("Rendered a Console Dialog")
}
