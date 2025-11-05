package main

import (
	"fmt"

	"example.com/complexproblems/gameengine"
)

func main() {
	// Test Case 1: Create a PC scene with buttons, textboxes, and particle effects
	fmt.Println("=== Test Case 1: PC Scene ===")
	pcFactory, err := gameengine.NewPlatformFactory("pc")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pcEngine := pcFactory.NewEngineBuilder().
		EnableParticleEffects().
		Build()

	pcSvc := pcFactory.Service()
	pcButton := pcSvc.CreateButton()
	pcTextbox := pcSvc.CreateTextbox()
	pcDialog := pcSvc.CreateDialog()

	pcButton.Render()
	pcTextbox.Render()
	pcDialog.Render()
	fmt.Printf("PC Engine - Shadows: %v, ParticleEffects: %v, Sound: %v\n",
		pcEngine.Shadows, pcEngine.ParticleEffects, pcEngine.Sound)

	// Test Case 2: Create a Mobile scene without shadows but with sounds
	fmt.Println("\n=== Test Case 2: Mobile Scene ===")
	mobileFactory, err := gameengine.NewPlatformFactory("mobile")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	mobileEngine := mobileFactory.NewEngineBuilder().
		EnableSound().
		Build()

	mobileSvc := mobileFactory.Service()
	mobileButton := mobileSvc.CreateButton()
	mobileTextbox := mobileSvc.CreateTextbox()
	mobileDialog := mobileSvc.CreateDialog()

	mobileButton.Render()
	mobileTextbox.Render()
	mobileDialog.Render()
	fmt.Printf("Mobile Engine - Shadows: %v, ParticleEffects: %v, Sound: %v\n",
		mobileEngine.Shadows, mobileEngine.ParticleEffects, mobileEngine.Sound)

	// Test Case 3: Verify that only one Physics Engine instance exists across multiple scenes
	fmt.Println("\n=== Test Case 3: Singleton Verification ===")
	consoleFactory, err := gameengine.NewPlatformFactory("console")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	engine1 := consoleFactory.NewEngineBuilder().
		EnableShadows().
		EnableSound().
		EnableParticleEffects().
		Build()

	engine2 := consoleFactory.NewEngineBuilder().
		EnableShadows().
		EnableSound().
		EnableParticleEffects().
		Build()

	// Also get the instance directly to verify
	engine3 := gameengine.GetInstance()

	fmt.Printf("engine1 == engine2: %v (should be true)\n", engine1 == engine2)
	fmt.Printf("engine1 == engine3: %v (should be true)\n", engine1 == engine3)
	fmt.Printf("engine2 == engine3: %v (should be true)\n", engine2 == engine3)
	fmt.Printf("Same Physics engine throughout the game: %v\n", engine1 == engine2)
}
