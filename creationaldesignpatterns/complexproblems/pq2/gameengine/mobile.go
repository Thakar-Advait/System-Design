package gameengine

import "fmt"

// MobileBuilder implements Builder for Mobile platform
type MobileBuilder struct {
	Shadows         bool
	ParticleEffects bool
	Sound           bool
	Platform        string
}

func (mobile *MobileBuilder) EnableShadows() Builder {
	mobile.Shadows = true
	return mobile
}

func (mobile *MobileBuilder) EnableParticleEffects() Builder {
	mobile.ParticleEffects = true
	return mobile
}

func (mobile *MobileBuilder) EnableSound() Builder {
	mobile.Sound = true
	return mobile
}

func (mobile *MobileBuilder) Build() *PhysicsEngine {
	engine := GetInstance()
	engine.Shadows = mobile.Shadows
	engine.ParticleEffects = mobile.ParticleEffects
	engine.Sound = mobile.Sound
	engine.Platform = mobile.Platform
	return engine
}

// MobileFactory implements Factory for Mobile platform
type MobileFactory struct{}

func (*MobileFactory) NewEngineBuilder() Builder {
	return &MobileBuilder{
		Platform: "mobile",
	}
}

func (*MobileFactory) Service() PlatformSvc {
	return &MobileSvc{}
}

// MobileSvc implements PlatformSvc for Mobile platform UI components
type MobileSvc struct{}

func (svc *MobileSvc) CreateButton() Button {
	return &MobileButton{}
}

func (svc *MobileSvc) CreateTextbox() Textbox {
	return &MobileTextbox{}
}

func (svc *MobileSvc) CreateDialog() Dialog {
	return &MobileDialog{}
}

// Mobile UI Component implementations
type MobileButton struct{}

func (*MobileButton) Render() {
	fmt.Println("Rendered a Mobile Button")
}

type MobileTextbox struct{}

func (*MobileTextbox) Render() {
	fmt.Println("Rendered a Mobile Textbox")
}

type MobileDialog struct{}

func (*MobileDialog) Render() {
	fmt.Println("Rendered a Mobile Dialog")
}
