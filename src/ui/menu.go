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
		fmt.Println("1 Info\n 2 Sac\n 3 Combat\n 4 Marchand\n 5 Forgeron\n 0 Quitter")
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
			p = bag.UseBag(p)
			fmt.Println("Entrée pour continuer")
			fmt.Scanln(&x)
		}
		if x == 3 {
			p = combat.Fight(p)
			fmt.Println("Entrée pour continuer")
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
