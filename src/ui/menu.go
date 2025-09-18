package ui

import (
	"fmt"
	"projetred/src/bag"
	"projetred/src/combat"
	"projetred/src/data"
	"projetred/src/forge"
	"projetred/src/merchant"
	"projetred/src/player"
)

func Menu(p data.Person) data.Person {
	for {
		fmt.Print(`
┌───────────────────────────────┐
│             MENU              │
├───────────────────────────────┤
│  [1] Info                     │
│  [2] Sac                      │
│  [3] Combat                   │
│  [4] Marchand                 │
│  [5] Forgeron                 │
│  [0] Quitter                  │
└───────────────────────────────┘
	Entrez votre choix :
`)
		var x int
		fmt.Scanln(&x)
		if x == 0 {
			return p
		}
		if x == 1 {
			player.Show(p)
			fmt.Println("Entrée pour continuer")
			fmt.Scanln(&x)
		}
		if x == 2 {
			p = bag.UseBag(p, data.Mob{})
			fmt.Println("\nEntrée pour continuer")
			fmt.Scanln(&x)
		}
		if x == 3 {
			p = combat.Fight(p)
			fmt.Println("\nEntrée pour continuer")
			fmt.Scanln(&x)
		}
		if x == 4 {
			p = merchant.Shop(p)
		}
		if x == 5 {
			p = forge.Smith(p)
		}
	}
}
