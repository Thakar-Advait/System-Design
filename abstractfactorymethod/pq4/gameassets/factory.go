package gameassets

import "fmt"

type Assets interface {
	CreateWeapon() Weapon
	CreateArmor() Armor
}

func NewAssetsFactory(vertical string) (Assets, error) {
	switch vertical {
	case "sci-fi":
		return SciFiFactory{}, nil
	case "medieval":
		return MedievalFactory{}, nil
	default:
		return nil, fmt.Errorf("unsupported asset type")
	}
}
