package gameengine

import "sync"

// PhysicsEngine represents the singleton physics engine instance
type PhysicsEngine struct {
	Shadows         bool
	ParticleEffects bool
	Sound           bool
	Platform        string
}

var (
	physicsEngineInstance *PhysicsEngine
	once                  sync.Once
)

// GetInstance returns the singleton instance of PhysicsEngine
func GetInstance() *PhysicsEngine {
	once.Do(func() {
		physicsEngineInstance = &PhysicsEngine{}
	})
	return physicsEngineInstance
}

// Builder interface for configuring game scenes with optional features
type Builder interface {
	EnableShadows() Builder
	EnableParticleEffects() Builder
	EnableSound() Builder
	Build() *PhysicsEngine
}

