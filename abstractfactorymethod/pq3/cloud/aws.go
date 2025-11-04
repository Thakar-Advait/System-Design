package cloud

type AWSComputeService struct{}
type AWSStorageService struct{}

func (AWSComputeService) Start() string {
	return "AWS Compute Service started"
}

func (AWSStorageService) Start() string {
	return "AWS Storage Service"
}

type AWSFactory struct{}

func (AWSFactory) CreateComputeService() ComputeService {
	return AWSComputeService{}
}

func (AWSFactory) CreateStorageService() StorageService {
	return AWSStorageService{}
}
