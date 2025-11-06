package main

import (
	"fmt"
	"sync"

	cloud "example.com/complexproblems/cloud"
)

func main() {
	fmt.Println("=== Testing Cloud Factory Implementation ===")

	// Test 1: Create factories for different cloud providers
	fmt.Println("Test 1: Creating factories for different cloud providers")
	awsFactory, err := cloud.NewCloudFactory("aws")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("✓ AWS factory created")

	azureFactory, err := cloud.NewCloudFactory("azure")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("✓ Azure factory created")

	gcpFactory, err := cloud.NewCloudFactory("gcp")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("✓ GCP factory created")

	// Test 2: Test singleton behavior - Service() should return same instance
	fmt.Println("\nTest 2: Verifying singleton behavior (if sync.Once is implemented)")
	awsService1 := awsFactory.Service()
	awsService2 := awsFactory.Service()
	fmt.Printf("  AWS Service 1: %p\n", awsService1)
	fmt.Printf("  AWS Service 2: %p\n", awsService2)

	// Test 3: Test builder pattern to create ResourceManager
	fmt.Println("\nTest 3: Building ResourceManager using builder pattern")

	// AWS Resource Manager
	awsResource := awsFactory.NewResourceManager().
		AlotCPU(4).
		AlotMemory(16).
		AlotStorage(100).
		AddTags("production").
		AddTags("web-server").
		Build()
	fmt.Printf("✓ AWS ResourceManager created: %+v\n", awsResource)

	// Azure Resource Manager
	azureResource := azureFactory.NewResourceManager().
		AlotCPU(8).
		AlotMemory(32).
		AlotStorage(200).
		AddTags("development").
		AddTags("database").
		Build()
	fmt.Printf("✓ Azure ResourceManager created: %+v\n", azureResource)

	// GCP Resource Manager
	gcpResource := gcpFactory.NewResourceManager().
		AlotCPU(2).
		AlotMemory(8).
		AlotStorage(50).
		AddTags("staging").
		AddTags("api-server").
		Build()
	fmt.Printf("✓ GCP ResourceManager created: %+v\n", gcpResource)

	// Test 4: Test Compute services
	fmt.Println("\nTest 4: Testing Compute services (VM deployment)")
	awsCloud := awsFactory.Service()
	awsCompute := awsCloud.SpawnCompute()
	awsVMResult := awsCompute.DeployVM(awsResource)
	fmt.Printf("✓ %s\n", awsVMResult)

	azureCloud := azureFactory.Service()
	azureCompute := azureCloud.SpawnCompute()
	azureVMResult := azureCompute.DeployVM(azureResource)
	fmt.Printf("✓ %s\n", azureVMResult)

	gcpCloud := gcpFactory.Service()
	gcpCompute := gcpCloud.SpawnCompute()
	gcpVMResult := gcpCompute.DeployVM(gcpResource)
	fmt.Printf("✓ %s\n", gcpVMResult)

	// Test 5: Test Storage services
	fmt.Println("\nTest 5: Testing Storage services")
	awsStorage := awsCloud.SpawnStorage()
	awsStoreResult := awsStorage.Store(awsResource, "user-123", "John Doe")
	fmt.Printf("✓ %s\n", awsStoreResult)

	azureStorage := azureCloud.SpawnStorage()
	azureStoreResult := azureStorage.Store(azureResource, "product-456", "Laptop")
	fmt.Printf("✓ %s\n", azureStoreResult)

	gcpStorage := gcpCloud.SpawnStorage()
	gcpStoreResult := gcpStorage.Store(gcpResource, "config-789", "api-key")
	fmt.Printf("✓ %s\n", gcpStoreResult)

	// Test 6: Test Network services
	fmt.Println("\nTest 6: Testing Network services (VPC deployment)")
	awsNetwork := awsCloud.SpawnNetwork()
	awsVPCResult := awsNetwork.DeployVPC(awsResource)
	fmt.Printf("✓ %s\n", awsVPCResult)

	azureNetwork := azureCloud.SpawnNetwork()
	azureVPCResult := azureNetwork.DeployVPC(azureResource)
	fmt.Printf("✓ %s\n", azureVPCResult)

	gcpNetwork := gcpCloud.SpawnNetwork()
	gcpVPCResult := gcpNetwork.DeployVPC(gcpResource)
	fmt.Printf("✓ %s\n", gcpVPCResult)

	// Test 7: Test complete workflow
	fmt.Println("\nTest 7: Testing complete workflow (Resource -> Compute -> Storage -> Network)")
	testResource := awsFactory.NewResourceManager().
		AlotCPU(16).
		AlotMemory(64).
		AlotStorage(500).
		AddTags("enterprise").
		AddTags("high-availability").
		Build()

	fmt.Printf("  Created resource: %+v\n", testResource)

	compute := awsCloud.SpawnCompute()
	storage := awsCloud.SpawnStorage()
	network := awsCloud.SpawnNetwork()

	fmt.Printf("  ✓ %s\n", compute.DeployVM(testResource))
	fmt.Printf("  ✓ %s\n", storage.Store(testResource, "data-key", "sensitive-data"))
	fmt.Printf("  ✓ %s\n", network.DeployVPC(testResource))

	// Test 8: Test concurrent access (if sync.Once is implemented)
	fmt.Println("\nTest 8: Testing concurrent access (thread safety)")
	var wg sync.WaitGroup
	instances := make([]cloud.Cloud, 100)

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(idx int) {
			defer wg.Done()
			instances[idx] = awsFactory.Service()
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
		fmt.Println("⚠ Concurrent calls returned different instances (sync.Once not implemented)")
	}

	// Test 9: Test error handling
	fmt.Println("\nTest 9: Testing error handling")
	invalidFactory, err := cloud.NewCloudFactory("invalid")
	if err != nil {
		fmt.Printf("✓ Error correctly returned for invalid cloud provider: %v\n", err)
	} else {
		fmt.Println("✗ Should have returned error for invalid cloud provider")
	}
	if invalidFactory != nil {
		fmt.Println("✗ Factory should be nil for invalid cloud provider")
	}

	// Test 10: Test builder flexibility
	fmt.Println("\nTest 10: Testing builder pattern flexibility")
	minimalResource := awsFactory.NewResourceManager().Build()
	fmt.Printf("✓ Minimal resource (defaults): %+v\n", minimalResource)

	fullResource := awsFactory.NewResourceManager().
		AlotCPU(32).
		AlotMemory(128).
		AlotStorage(1000).
		AddTags("tag1").
		AddTags("tag2").
		AddTags("tag3").
		Build()
	fmt.Printf("✓ Full resource (all fields): %+v\n", fullResource)

	// Test 11: Test Global Resource Manager singleton with sync.Once
	fmt.Println("\nTest 11: Testing Global Resource Manager singleton (sync.Once)")
	manager1 := cloud.GetGlobalResourceManager()
	manager2 := cloud.GetGlobalResourceManager()
	manager3 := cloud.GetGlobalResourceManager()

	if manager1 == manager2 && manager2 == manager3 {
		fmt.Printf("✓ All calls returned the same instance (address: %p)\n", manager1)
	} else {
		fmt.Println("✗ Different instances returned!")
	}

	// Test concurrent access to verify sync.Once works correctly
	var wg2 sync.WaitGroup
	managers := make([]*cloud.GlobalResourceManager, 100)
	wg2.Add(100)
	for i := 0; i < 100; i++ {
		go func(idx int) {
			defer wg2.Done()
			managers[idx] = cloud.GetGlobalResourceManager()
		}(i)
	}
	wg2.Wait()

	allSameManager := true
	firstManager := managers[0]
	for i := 1; i < 100; i++ {
		if managers[i] != firstManager {
			allSameManager = false
			break
		}
	}

	if allSameManager {
		fmt.Printf("✓ All 100 concurrent calls returned the same instance (address: %p)\n", firstManager)
	} else {
		fmt.Println("✗ Concurrent calls returned different instances!")
	}

	// Test 12: Test deployment tracking
	fmt.Println("\nTest 12: Testing deployment tracking across all cloud providers")
	fmt.Printf("  Total deployments tracked: %d\n", manager1.GetTotalDeployments())

	// Get summary
	summary := manager1.GetSummary()
	fmt.Println("\n  Deployment Summary:")
	fmt.Println(summary)

	// Test filtering by cloud provider
	fmt.Println("\n  Deployments by Cloud Provider:")
	awsDeployments := manager1.GetDeploymentsByCloudProvider("AWS")
	fmt.Printf("    AWS: %d deployments\n", len(awsDeployments))
	azureDeployments := manager1.GetDeploymentsByCloudProvider("Azure")
	fmt.Printf("    Azure: %d deployments\n", len(azureDeployments))
	gcpDeployments := manager1.GetDeploymentsByCloudProvider("GCP")
	fmt.Printf("    GCP: %d deployments\n", len(gcpDeployments))

	// Test filtering by resource type
	fmt.Println("\n  Deployments by Resource Type:")
	vmDeployments := manager1.GetDeploymentsByResourceType("VM")
	fmt.Printf("    VM: %d deployments\n", len(vmDeployments))
	storageDeployments := manager1.GetDeploymentsByResourceType("Storage")
	fmt.Printf("    Storage: %d deployments\n", len(storageDeployments))
	vpcDeployments := manager1.GetDeploymentsByResourceType("VPC")
	fmt.Printf("    VPC: %d deployments\n", len(vpcDeployments))

	// Show a few sample deployments
	fmt.Println("\n  Sample Deployments:")
	allDeployments := manager1.GetDeployments()
	maxSamples := 3
	if len(allDeployments) < maxSamples {
		maxSamples = len(allDeployments)
	}
	for i := 0; i < maxSamples; i++ {
		dep := allDeployments[i]
		fmt.Printf("    [%s] %s %s - %s\n", dep.ID, dep.CloudProvider, dep.ResourceType, dep.Details)
	}

	fmt.Println("\n=== All tests completed ===")
}
