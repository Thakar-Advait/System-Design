package cloud

type Compute interface {
	DeployVM(resource ResourceManager) string
}

type Storage interface {
	Store(resource ResourceManager, key, value string) string
}

type Network interface {
	DeployVPC(resource ResourceManager) string
}
