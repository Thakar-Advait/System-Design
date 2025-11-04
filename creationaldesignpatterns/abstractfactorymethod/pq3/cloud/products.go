package cloud

type ComputeService interface {
	Start() string
}

type StorageService interface {
	Start() string
}