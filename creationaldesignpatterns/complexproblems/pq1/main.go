package main

import (
	"fmt"
	"time"

	"example.com/complexproblems/messaging"
)

func main() {
	// --- Email example ---
	emailFactory, err := messaging.NewPlatformFactory("email")
	if err != nil {
		panic(err)
	}

	emailBuilder := emailFactory.NewBuilder()
	emailMsg, err := emailBuilder.
		AddHeader("X-Custom: header1").
		AddFooter("Regards, Team").
		AddInlineImage("https://cdn.example.com/img1.png").
		SetBody("Hello Advait! This is an email body.").
		Build()
	if err != nil {
		panic(err)
	}

	emailService := emailFactory.Service()
	_ = emailService.Send(emailMsg)

	// schedule 2 seconds in future (demo)
	_ = emailService.Schedule(emailMsg, time.Now().Add(2*time.Second))

	// --- SMS example ---
	smsFactory, err := messaging.NewPlatformFactory("sms")
	if err != nil {
		panic(err)
	}

	smsBuilder := smsFactory.NewBuilder()
	smsMsg, err := smsBuilder.
		SetBody("Short SMS body").
		Build()
	if err != nil {
		panic(err)
	}

	smsService := smsFactory.Service()
	_ = smsService.Send(smsMsg)

	// Demonstrate that service instances are singletons:
	fmt.Printf("Email service same as second fetch: %v\n", emailFactory.Service() == emailService)
	fmt.Printf("SMS service same as second fetch: %v\n", smsFactory.Service() == smsService)

	// wait for scheduled send to print (only for demo)
	time.Sleep(3 * time.Second)
}
