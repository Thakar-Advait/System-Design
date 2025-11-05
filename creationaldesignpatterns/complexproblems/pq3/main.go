package main

import (
	"fmt"
	"sync"

	orderengine "example.com/complexproblems/order_engine"
)

func main() {
	fmt.Println("=== Testing Order Engine Implementation ===\n")

	// Test 1: Create factories for different platforms
	fmt.Println("Test 1: Creating factories for different platforms")
	stripeFactory, err := orderengine.NewOrderFactory("stripe")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("✓ Stripe factory created")

	paypalFactory, err := orderengine.NewOrderFactory("paypal")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("✓ PayPal factory created")

	bankFactory, err := orderengine.NewOrderFactory("bank")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("✓ Bank factory created")

	// Test 2: Test singleton behavior - Service() should return same instance
	fmt.Println("\nTest 2: Verifying singleton behavior (sync.Once)")
	stripeService1 := stripeFactory.Service()
	stripeService2 := stripeFactory.Service()
	if stripeService1 == stripeService2 {
		fmt.Printf("✓ Stripe service returns same instance (address: %p)\n", stripeService1)
	} else {
		fmt.Println("✗ Stripe service returns different instances!")
	}

	paypalService1 := paypalFactory.Service()
	paypalService2 := paypalFactory.Service()
	if paypalService1 == paypalService2 {
		fmt.Printf("✓ PayPal service returns same instance (address: %p)\n", paypalService1)
	} else {
		fmt.Println("✗ PayPal service returns different instances!")
	}

	// Test 3: Test builder pattern to create invoices
	fmt.Println("\nTest 3: Building invoices using builder pattern")

	stripeInvoice := stripeFactory.NewInvoiceBuilder().
		ApplyDiscount(10).
		AddTax(8).
		AddShippingInfo("Standard Shipping").
		GiftWrap().
		Build()
	fmt.Printf("✓ Stripe invoice created: %+v\n", stripeInvoice)

	paypalInvoice := paypalFactory.NewInvoiceBuilder().
		ApplyDiscount(15).
		AddTax(10).
		AddShippingInfo("Express Shipping").
		Build()
	fmt.Printf("✓ PayPal invoice created: %+v\n", paypalInvoice)

	bankInvoice := bankFactory.NewInvoiceBuilder().
		ApplyDiscount(5).
		AddTax(12).
		AddShippingInfo("Bank Transfer").
		GiftWrap().
		Build()
	fmt.Printf("✓ Bank invoice created: %+v\n", bankInvoice)

	// Test 4: Place orders through different platforms
	fmt.Println("\nTest 4: Placing orders through different platforms")
	stripeOrder := stripeFactory.Service().Normal(stripeInvoice)
	fmt.Printf("✓ %s\n", stripeOrder)

	paypalOrder := paypalFactory.Service().Express(paypalInvoice)
	fmt.Printf("✓ %s\n", paypalOrder)

	bankOrder := bankFactory.Service().Subscription(bankInvoice)
	fmt.Printf("✓ %s\n", bankOrder)

	// Test 5: Test concurrent access to verify thread safety
	fmt.Println("\nTest 5: Testing concurrent access (thread safety)")
	var wg sync.WaitGroup
	instances := make([]orderengine.Platform, 100)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(idx int) {
			defer wg.Done()
			instances[idx] = stripeFactory.Service()
		}(i)
	}
	wg.Wait()

	// Verify all instances are the same
	allSame := true
	firstInstance := instances[0]
	for i := 1; i < 100; i++ {
		if instances[i] != firstInstance {
			allSame = false
			break
		}
	}

	if allSame {
		fmt.Printf("✓ All 100 concurrent calls returned the same instance (address: %p)\n", firstInstance)
	} else {
		fmt.Println("✗ Concurrent calls returned different instances!")
	}

	// Test 6: Test error handling
	fmt.Println("\nTest 6: Testing error handling")
	invalidFactory, err := orderengine.NewOrderFactory("invalid")
	if err != nil {
		fmt.Printf("✓ Error correctly returned for invalid platform: %v\n", err)
	} else {
		fmt.Println("✗ Should have returned error for invalid platform")
	}
	if invalidFactory != nil {
		fmt.Println("✗ Factory should be nil for invalid platform")
	}

	fmt.Println("\n=== All tests completed ===")
}
