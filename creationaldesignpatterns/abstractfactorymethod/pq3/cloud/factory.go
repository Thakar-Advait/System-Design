package cloud

import "fmt"

type CloudFactory interface {
	CreateComputeService() ComputeService
	CreateStorageService() StorageService
}

func NewCloudFactory(vertical string) (CloudFactory, error) {
	switch vertical {
	case "aws":
		return AWSFactory{}, nil
	case "gcp":
		return GCPFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported cloud provider")
	}
}
