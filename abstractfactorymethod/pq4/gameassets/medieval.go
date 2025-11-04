package gameassets

type MedievalWeapon struct{}
type MedievalArmor struct{}

func (MedievalWeapon) Use() string {
	return "Mjolnir"
}

func (MedievalArmor) Protect() string {
	return "Knight's Guard"
}

type MedievalFactory struct{}

func (MedievalFactory) CreateWeapon() Weapon {
	return MedievalWeapon{}
}

func (MedievalFactory) CreateArmor() Armor {
	return MedievalArmor{}
}
