package cloud

import "fmt"

type GCP struct{}
type GCPFactory struct{}
type GCPSvc struct{}
type GCPResourceBuilder struct {
	cpu     int32
	memory  int32
	storage int32
	tags    []string
}

func (fac *GCPFactory) NewResourceManager() Builder {
	return &GCPResourceBuilder{
		tags: make([]string, 0),
	}
}

func (b *GCPResourceBuilder) AlotCPU(cpu int32) Builder {
	b.cpu = cpu
	return b
}

func (b *GCPResourceBuilder) AlotMemory(memory int32) Builder {
	b.memory = memory
	return b
}

func (b *GCPResourceBuilder) AlotStorage(storage int32) Builder {
	b.storage = storage
	return b
}

func (b *GCPResourceBuilder) AddTags(tag string) Builder {
	b.tags = append(b.tags, tag)
	return b
}

func (b *GCPResourceBuilder) Build() ResourceManager {
	return ResourceManager{
		CPU:     b.cpu,
		Memory:  b.memory,
		Storage: b.storage,
		Tags:    b.tags,
	}
}

type GCPCompute struct{}
type GCPStorage struct{}
type GCPNetwork struct{}

func (*GCP) SpawnCompute() Compute {
	return &GCPCompute{}
}

func (*GCPCompute) DeployVM(resource ResourceManager) string {
	manager := GetGlobalResourceManager()
	deploymentID := manager.RegisterDeployment("GCP", "VM", resource, fmt.Sprintf("GCP VM deployed with configurations: %v", resource))
	return fmt.Sprintf("GCP VM deployed with configurations: %v [ID: %s]", resource, deploymentID)
}

func (*GCP) SpawnStorage() Storage {
	return &GCPStorage{}
}

func (*GCPStorage) Store(resource ResourceManager, key, val string) string {
	manager := GetGlobalResourceManager()
	details := fmt.Sprintf("Key: %s Value: %s stored in GCP Store with configurations: %v", key, val, resource)
	deploymentID := manager.RegisterDeployment("GCP", "Storage", resource, details)
	return fmt.Sprintf("Key: %s Value: %s stored in GCP Store with configurations: %v [ID: %s]", key, val, resource, deploymentID)
}

func (*GCP) SpawnNetwork() Network {
	return &GCPNetwork{}
}

func (*GCPNetwork) DeployVPC(resource ResourceManager) string {
	manager := GetGlobalResourceManager()
	deploymentID := manager.RegisterDeployment("GCP", "VPC", resource, fmt.Sprintf("GCP VPC deployed with configurations: %v", resource))
	return fmt.Sprintf("GCP VPC deployed with configurations: %v [ID: %s]", resource, deploymentID)
}

func (*GCPFactory) Service() Cloud {
	return &GCP{}
}

func (*GCPSvc) CreateComputeService() Compute {
	return &GCPCompute{}
}

func (*GCPSvc) CreateStorageService() Storage {
	return &GCPStorage{}
}

func (*GCPSvc) CreateNetworkService() Network {
	return &GCPNetwork{}
}
