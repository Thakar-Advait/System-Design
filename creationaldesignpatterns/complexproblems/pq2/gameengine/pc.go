package gameengine

import "fmt"

// PCBuilder implements Builder for PC platform
type PCBuilder struct {
	Shadows         bool
	ParticleEffects bool
	Sound           bool
	Platform        string
}

func (pc *PCBuilder) EnableShadows() Builder {
	pc.Shadows = true
	return pc
}

func (pc *PCBuilder) EnableParticleEffects() Builder {
	pc.ParticleEffects = true
	return pc
}

func (pc *PCBuilder) EnableSound() Builder {
	pc.Sound = true
	return pc
}

func (pc *PCBuilder) Build() *PhysicsEngine {
	engine := GetInstance()
	engine.Shadows = pc.Shadows
	engine.ParticleEffects = pc.ParticleEffects
	engine.Sound = pc.Sound
	engine.Platform = pc.Platform
	return engine
}

// PCFactory implements Factory for PC platform
type PCFactory struct{}

func (*PCFactory) NewEngineBuilder() Builder {
	return &PCBuilder{
		Platform: "pc",
	}
}

func (*PCFactory) Service() PlatformSvc {
	return &PCSvc{}
}

// PCSvc implements PlatformSvc for PC platform UI components
type PCSvc struct{}

func (svc *PCSvc) CreateButton() Button {
	return &PCButton{}
}

func (svc *PCSvc) CreateTextbox() Textbox {
	return &PCTextbox{}
}

func (svc *PCSvc) CreateDialog() Dialog {
	return &PCDialog{}
}

// PC UI Component implementations
type PCButton struct{}

func (*PCButton) Render() {
	fmt.Println("Rendered a PC Button")
}

type PCTextbox struct{}

func (*PCTextbox) Render() {
	fmt.Println("Rendered a PC Textbox")
}

type PCDialog struct{}

func (*PCDialog) Render() {
	fmt.Println("Rendered a PC Dialog")
}
