package gameassets

type SciFiWeapon struct{}
type SciFiArmor struct{}

func (SciFiWeapon) Use() string {
	return "Plasma Gun"
}

func (SciFiArmor) Protect() string {
	return "Nanotech Armor"
}

type SciFiFactory struct{}

func (SciFiFactory) CreateWeapon() Weapon {
	return SciFiWeapon{}
}

func (SciFiFactory) CreateArmor() Armor {
	return SciFiArmor{}
}
