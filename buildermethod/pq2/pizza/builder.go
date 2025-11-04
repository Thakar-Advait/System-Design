package pizza

type pizza struct {
	base     string
	size     string
	toppings []string
}

type Builder interface {
	SetBase(base string) Builder
	SetSize(size string) Builder
	AddTopping(topping string) Builder
	Build() pizza
}

type pizzaBuilder struct {
	base     string
	size     string
	toppings []string
}

func NewBuilder() Builder {
	return &pizzaBuilder{}
}

func (p *pizzaBuilder) SetBase(base string) Builder {
	p.base = base
	return p
	// panic("to be implemented")
}

func (p *pizzaBuilder) SetSize(size string) Builder {
	p.size = size
	return p
	// panic("to be implemented")
}

func (p *pizzaBuilder) AddTopping(topping string) Builder {
	p.toppings = append(p.toppings, topping)
	return p
	// panic("to be implemented")
}

func (p *pizzaBuilder) Build() pizza {
	return pizza{
		base:     p.base,
		size:     p.size,
		toppings: p.toppings,
	}
	// panic("to be implemented")
}
