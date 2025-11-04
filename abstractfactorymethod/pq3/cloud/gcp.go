package cloud

type GCPComputeService struct{}
type GCPStorageService struct{}

func (GCPComputeService) Start() string {
	return "GCP Compute Service started"
}

func (GCPStorageService) Start() string {
	return "GCP Storage Service started"
}

type GCPFactory struct{}

func (GCPFactory) CreateComputeService() ComputeService {
	return GCPComputeService{}
}

func (GCPFactory) CreateStorageService() StorageService {
	return GCPStorageService{}
}
