package main

import (
	"fmt"

	"example.com/buildermethod/resume"
)

func main() {
	res := resume.NewResumeBuilder().
		SetName("Advait Thakar").
		AddEducation("B.Tech").
		AddExperience("Cognam Technologies").
		Build()
	fmt.Printf("Resume: %v\n", res)
}
