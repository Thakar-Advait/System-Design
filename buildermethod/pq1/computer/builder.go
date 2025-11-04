package computer

type Builder interface {
	SetCPU(config string) Builder
	SetRAM(config string) Builder
	SetStorage(config string) Builder
	SetGPU(config string) Builder
	Build() Computer
}

type computerBuilder struct {
	cpu     string
	ram     string
	storage string
	gpu     string
}

func NewBuilder() Builder {
	return &computerBuilder{}
}

func (c *computerBuilder) SetCPU(cpu string) Builder {
	c.cpu = cpu
	return c
}

func (c *computerBuilder) SetGPU(gpu string) Builder {
	c.gpu = gpu
	return c
}

func (c *computerBuilder) SetStorage(storage string) Builder {
	c.storage = storage
	return c
}

func (c *computerBuilder) SetRAM(ram string) Builder {
	c.ram = ram
	return c
}

func (c *computerBuilder) Build() Computer {
	return Computer{
		CPU:     c.cpu,
		GPU:     c.gpu,
		RAM:     c.ram,
		Storage: c.storage,
	}
}

