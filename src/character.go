package main

import (
	"fmt"
)

type Character struct {
	Name      string
	Classe    string
	Lvl       int
	Inventory []string
	Pvmax     int
	Pvactuel  int
}

func CharacterInfo(Name string, Classe string, Lvl int, Inventory []string, Pvmax int, Pvactuel int) {
	Character := Character{
		Name:      Name,
		Classe:    Classe,
		Lvl:       Lvl,
		Pvmax:     Pvmax,
		Pvactuel:  Pvactuel,
		Inventory: []string{},
	}

	fmt.Println("Nom", Character.Name)
	fmt.Println("Classe", Character.Classe)
	fmt.Println("Lvl", Character.Lvl)
	fmt.Println("Pvmax", Character.Pvmax)
	fmt.Println("Pvactuel", Character.Pvactuel)
	fmt.Println("inventory", Character.Inventory)
}
