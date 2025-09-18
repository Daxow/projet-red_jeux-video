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
		fmt.Println(`
┌───────────────────────────────┐
│    CHOISISSEZ VOTRE CLASSE    │
├───────────────────────────────┤
│  [1] Humain                   │
│      ───────                  │
│                               │
│  [2] Elfe                     │
│      ───────                  │
│                               │
│  [3] Nain                     │
│      ───────                  │
└───────────────────────────────┘
	Entrez votre choix :`)
		fmt.Scanln(&c)
		if c == "1" || c == "2" || c == "3" {
			break
		}
	}
	pv := 100
	if c == "1" {
		c = "Humain"
	}
	if c == "2" {
		pv = 80
		c = "Elfe"
	}
	if c == "3" {
		pv = 120
		c = "Nain"
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
	p.Size = 10
	p.Up = 0
	return p
}

func Show(p data.Person) {
	fmt.Println("\nNom:", p.Name)
	fmt.Println("Classe:", p.Class)
	fmt.Println("Niveau:", p.Level)
	fmt.Println("PV:", p.Hp, "/", p.Hpmax)
	fmt.Println("Or:", p.Gold)
	fmt.Println("Sorts:", p.Spells)
	fmt.Println("Equipements:", p.Eq)
	fmt.Print ("\n")
}

func Dead(p data.Person) data.Person {
	if p.Hp <= 0 {
		p.Hp = p.Hpmax / 2
		fmt.Println("Mort, retour avec", p.Hp)
	}
	return p
}
