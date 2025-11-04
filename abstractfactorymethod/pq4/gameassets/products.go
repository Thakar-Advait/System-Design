package gameassets

type Weapon interface {
	Use() string
}

type Armor interface {
	Protect() string
}
