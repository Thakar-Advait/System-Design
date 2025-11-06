package cloud

import "fmt"

type AWS struct{}
type AWSFactory struct{}
type AWSSvc struct{}
type AWSResourceBuilder struct {
	cpu     int32
	memory  int32
	storage int32
	tags    []string
}

func (fac *AWSFactory) NewResourceManager() Builder {
	return &AWSResourceBuilder{
		tags: make([]string, 0),
	}
}

func (b *AWSResourceBuilder) AlotCPU(cpu int32) Builder {
	b.cpu = cpu
	return b
}

func (b *AWSResourceBuilder) AlotMemory(memory int32) Builder {
	b.memory = memory
	return b
}

func (b *AWSResourceBuilder) AlotStorage(storage int32) Builder {
	b.storage = storage
	return b
}

func (b *AWSResourceBuilder) AddTags(tag string) Builder {
	b.tags = append(b.tags, tag)
	return b
}

func (b *AWSResourceBuilder) Build() ResourceManager {
	return ResourceManager{
		CPU:     b.cpu,
		Memory:  b.memory,
		Storage: b.storage,
		Tags:    b.tags,
	}
}

type AWSCompute struct{}
type AWSStorage struct{}
type AWSNetwork struct{}

func (*AWS) SpawnCompute() Compute {
	return &AWSCompute{}
}

func (*AWSCompute) DeployVM(resource ResourceManager) string {
	manager := GetGlobalResourceManager()
	deploymentID := manager.RegisterDeployment("AWS", "VM", resource, fmt.Sprintf("AWS VM deployed with configurations: %v", resource))
	return fmt.Sprintf("AWS VM deployed with configurations: %v [ID: %s]", resource, deploymentID)
}

func (*AWS) SpawnStorage() Storage {
	return &AWSStorage{}
}

func (*AWSStorage) Store(resource ResourceManager, key, val string) string {
	manager := GetGlobalResourceManager()
	details := fmt.Sprintf("Key: %s Value: %s stored in AWS Store with configurations: %v", key, val, resource)
	deploymentID := manager.RegisterDeployment("AWS", "Storage", resource, details)
	return fmt.Sprintf("Key: %s Value: %s stored in AWS Store with configurations: %v [ID: %s]", key, val, resource, deploymentID)
}

func (*AWS) SpawnNetwork() Network {
	return &AWSNetwork{}
}

func (*AWSNetwork) DeployVPC(resource ResourceManager) string {
	manager := GetGlobalResourceManager()
	deploymentID := manager.RegisterDeployment("AWS", "VPC", resource, fmt.Sprintf("AWS VPC deployed with configurations: %v", resource))
	return fmt.Sprintf("AWS VPC deployed with configurations: %v [ID: %s]", resource, deploymentID)
}

func (*AWSFactory) Service() Cloud {
	return &AWS{}
}

func (*AWSSvc) CreateComputeService() Compute {
	return &AWSCompute{}
}

func (*AWSSvc) CreateStorageService() Storage {
	return &AWSStorage{}
}

func (*AWSSvc) CreateNetworkService() Network {
	return &AWSNetwork{}
}
