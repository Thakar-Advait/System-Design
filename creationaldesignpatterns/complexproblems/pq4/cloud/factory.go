package cloud

import "fmt"

type Cloud interface {
	SpawnCompute() Compute
	SpawnStorage() Storage
	SpawnNetwork() Network
}

type ResourceManager struct {
	CPU     int32
	Memory  int32
	Storage int32
	Tags    []string
}

type Builder interface {
	AlotCPU(cpu int32) Builder
	AlotMemory(memory int32) Builder
	AlotStorage(storage int32) Builder
	AddTags(tag string) Builder
	Build() ResourceManager
}

type Factory interface {
	NewResourceManager() Builder
	Service() Cloud
}

func NewCloudFactory(vertical string) (Factory, error) {
	switch vertical {
	case "aws":
		return &AWSFactory{}, nil
	case "gcp":
		return &GCPFactory{}, nil
	case "azure":
		return &AzureFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported cloud")
	}
}
