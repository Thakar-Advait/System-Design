package cloud

import "fmt"

type Azure struct{}
type AzureFactory struct{}
type AzureSvc struct{}
type AzureResourceBuilder struct {
	cpu     int32
	memory  int32
	storage int32
	tags    []string
}

func (fac *AzureFactory) NewResourceManager() Builder {
	return &AzureResourceBuilder{
		tags: make([]string, 0),
	}
}

func (b *AzureResourceBuilder) AlotCPU(cpu int32) Builder {
	b.cpu = cpu
	return b
}

func (b *AzureResourceBuilder) AlotMemory(memory int32) Builder {
	b.memory = memory
	return b
}

func (b *AzureResourceBuilder) AlotStorage(storage int32) Builder {
	b.storage = storage
	return b
}

func (b *AzureResourceBuilder) AddTags(tag string) Builder {
	b.tags = append(b.tags, tag)
	return b
}

func (b *AzureResourceBuilder) Build() ResourceManager {
	return ResourceManager{
		CPU:     b.cpu,
		Memory:  b.memory,
		Storage: b.storage,
		Tags:    b.tags,
	}
}

type AzureCompute struct{}
type AzureStorage struct{}
type AzureNetwork struct{}

func (*Azure) SpawnCompute() Compute {
	return &AzureCompute{}
}

func (*AzureCompute) DeployVM(resource ResourceManager) string {
	manager := GetGlobalResourceManager()
	deploymentID := manager.RegisterDeployment("Azure", "VM", resource, fmt.Sprintf("Azure VM deployed with configurations: %v", resource))
	return fmt.Sprintf("Azure VM deployed with configurations: %v [ID: %s]", resource, deploymentID)
}

func (*Azure) SpawnStorage() Storage {
	return &AzureStorage{}
}

func (*AzureStorage) Store(resource ResourceManager, key, val string) string {
	manager := GetGlobalResourceManager()
	details := fmt.Sprintf("Key: %s Value: %s stored in Azure Store with configurations: %v", key, val, resource)
	deploymentID := manager.RegisterDeployment("Azure", "Storage", resource, details)
	return fmt.Sprintf("Key: %s Value: %s stored in Azure Store with configurations: %v [ID: %s]", key, val, resource, deploymentID)
}

func (*Azure) SpawnNetwork() Network {
	return &AzureNetwork{}
}

func (*AzureNetwork) DeployVPC(resource ResourceManager) string {
	manager := GetGlobalResourceManager()
	deploymentID := manager.RegisterDeployment("Azure", "VPC", resource, fmt.Sprintf("Azure VPC deployed with configurations: %v", resource))
	return fmt.Sprintf("Azure VPC deployed with configurations: %v [ID: %s]", resource, deploymentID)
}

func (*AzureFactory) Service() Cloud {
	return &Azure{}
}

func (*AzureSvc) CreateComputeService() Compute {
	return &AzureCompute{}
}

func (*AzureSvc) CreateStorageService() Storage {
	return &AzureStorage{}
}

func (*AzureSvc) CreateNetworkService() Network {
	return &AzureNetwork{}
}
