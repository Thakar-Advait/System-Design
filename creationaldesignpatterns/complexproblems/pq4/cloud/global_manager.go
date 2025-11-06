package cloud

import (
	"fmt"
	"sync"
)

// Deployment represents a tracked deployment
type Deployment struct {
	ID            string
	CloudProvider string
	ResourceType  string // "VM", "Storage", "VPC"
	Resource      ResourceManager
	Details       string
}

// GlobalResourceManager is a singleton that tracks all deployments across all cloud providers
type GlobalResourceManager struct {
	deployments []Deployment
	mu          sync.RWMutex
	nextID      int
}

var (
	globalManagerInstance *GlobalResourceManager
	once                  sync.Once
)

// GetGlobalResourceManager returns the singleton instance of GlobalResourceManager
// Uses sync.Once to ensure only one instance is created, even with concurrent access
func GetGlobalResourceManager() *GlobalResourceManager {
	once.Do(func() {
		globalManagerInstance = &GlobalResourceManager{
			deployments: make([]Deployment, 0),
			nextID:      1,
		}
	})
	return globalManagerInstance
}

// RegisterDeployment registers a new deployment in the global manager
func (grm *GlobalResourceManager) RegisterDeployment(cloudProvider, resourceType string, resource ResourceManager, details string) string {
	grm.mu.Lock()
	defer grm.mu.Unlock()

	deploymentID := fmt.Sprintf("%s-%s-%d", cloudProvider, resourceType, grm.nextID)
	grm.nextID++

	deployment := Deployment{
		ID:            deploymentID,
		CloudProvider: cloudProvider,
		ResourceType:  resourceType,
		Resource:      resource,
		Details:       details,
	}

	grm.deployments = append(grm.deployments, deployment)
	return deploymentID
}

// GetDeployments returns all tracked deployments
func (grm *GlobalResourceManager) GetDeployments() []Deployment {
	grm.mu.RLock()
	defer grm.mu.RUnlock()

	// Return a copy to prevent external modifications
	deployments := make([]Deployment, len(grm.deployments))
	copy(deployments, grm.deployments)
	return deployments
}

// GetDeploymentsByCloudProvider returns deployments filtered by cloud provider
func (grm *GlobalResourceManager) GetDeploymentsByCloudProvider(cloudProvider string) []Deployment {
	grm.mu.RLock()
	defer grm.mu.RUnlock()

	var filtered []Deployment
	for _, deployment := range grm.deployments {
		if deployment.CloudProvider == cloudProvider {
			filtered = append(filtered, deployment)
		}
	}
	return filtered
}

// GetDeploymentsByResourceType returns deployments filtered by resource type
func (grm *GlobalResourceManager) GetDeploymentsByResourceType(resourceType string) []Deployment {
	grm.mu.RLock()
	defer grm.mu.RUnlock()

	var filtered []Deployment
	for _, deployment := range grm.deployments {
		if deployment.ResourceType == resourceType {
			filtered = append(filtered, deployment)
		}
	}
	return filtered
}

// GetTotalDeployments returns the total number of deployments
func (grm *GlobalResourceManager) GetTotalDeployments() int {
	grm.mu.RLock()
	defer grm.mu.RUnlock()
	return len(grm.deployments)
}

// GetSummary returns a summary of all deployments
func (grm *GlobalResourceManager) GetSummary() string {
	grm.mu.RLock()
	defer grm.mu.RUnlock()

	if len(grm.deployments) == 0 {
		return "No deployments tracked yet."
	}

	summary := fmt.Sprintf("Total Deployments: %d\n", len(grm.deployments))

	// Count by cloud provider
	cloudCounts := make(map[string]int)
	resourceCounts := make(map[string]int)

	for _, deployment := range grm.deployments {
		cloudCounts[deployment.CloudProvider]++
		resourceCounts[deployment.ResourceType]++
	}

	summary += "\nBy Cloud Provider:\n"
	for cloud, count := range cloudCounts {
		summary += fmt.Sprintf("  %s: %d\n", cloud, count)
	}

	summary += "\nBy Resource Type:\n"
	for resourceType, count := range resourceCounts {
		summary += fmt.Sprintf("  %s: %d\n", resourceType, count)
	}

	return summary
}
