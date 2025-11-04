package main

import (
	"fmt"

	"example.com/gameassets/gameassets"
)

func main() {
	assets, err := gameassets.NewAssetsFactory("medieval")
	if err != nil {
		fmt.Println(err.Error())
	}
	weapon := assets.CreateWeapon()
	armor := assets.CreateArmor()

	fmt.Println(weapon.Use())
	fmt.Println(armor.Protect())
}
