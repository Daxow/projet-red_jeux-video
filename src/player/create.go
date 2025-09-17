package player

import (
	"fmt"
	"projetred/src/data"
)

func CreateCharacter() data.Person {
	var n string
	var c string
	for {
		fmt.Println("Nom ?")
		fmt.Scanln(&n)
		if len(n) > 0 {
			break
		}
	}
	for {
		fmt.Println("Classe ? Humain / Elfe / Nain")
		fmt.Scanln(&c)
		if c == "Humain" || c == "Elfe" || c == "Nain" {
			break
		}
	}
	pv := 100
	if c == "Elfe" {
		pv = 80
	}
	if c == "Nain" {
		pv = 120
	}

	p := data.Person{}
	p.Name = n
	p.Class = c
	p.Level = 1
	p.Hpmax = pv
	p.Hp = pv / 2
	p.Gold = 100
	p.Spells = []string{"Coup de poing"}
	p.Bag = []string{"Potion de vie", "Potion de vie"}
	p.Size = 5
	p.Up = 0
	return p
}

func Show(p data.Person) {
	fmt.Println("Nom:", p.Name)
	fmt.Println("Classe:", p.Class)
	fmt.Println("Niveau:", p.Level)
	fmt.Println("PV:", p.Hp, "/", p.Hpmax)
	fmt.Println("Or:", p.Gold)
	fmt.Println("Sorts:", p.Spells)
	fmt.Println("Equipements:", p.Eq)
}

func Dead(p data.Person) data.Person {
	if p.Hp <= 0 {
		p.Hp = p.Hpmax / 2
		fmt.Println("Mort, retour avec", p.Hp)
	}
	return p
}
