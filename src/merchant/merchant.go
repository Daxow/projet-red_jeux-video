package merchant

import (
	"fmt"
	"projetred/src/bag"
	"projetred/src/data"
)

func Shop(p data.Person) data.Person {
	for {
		fmt.Println("Marchand:")
		fmt.Println("1. Potion de vie (3)")
		fmt.Println("2. Potion de poison (6)")
		fmt.Println("3. Livre de Sort : Boule de Feu (25)")
		fmt.Println("4. Fourrure de Loup (4)")
		fmt.Println("5. Peau de Troll (7)")
		fmt.Println("6. Cuir de Sanglier (3)")
		fmt.Println("7. Plume de Corbeau (1)")
		fmt.Println("8. Augmentation d’inventaire (30)")
		fmt.Println("0. Retour")
		var x int
		fmt.Scanln(&x)
		if x == 0 {
			return p
		}

		name := ""
		cost := 0
		if x == 1 {
			name = "Potion de vie"
			cost = 3
		}
		if x == 2 {
			name = "Potion de poison"
			cost = 6
		}
		if x == 3 {
			name = "Livre de Sort : Boule de Feu"
			cost = 25
		}
		if x == 4 {
			name = "Fourrure de Loup"
			cost = 4
		}
		if x == 5 {
			name = "Peau de Troll"
			cost = 7
		}
		if x == 6 {
			name = "Cuir de Sanglier"
			cost = 3
		}
		if x == 7 {
			name = "Plume de Corbeau"
			cost = 1
		}
		if x == 8 {
			name = "Augmentation d’inventaire"
			cost = 30
		}

		if name == "" {
			continue
		}
		if p.Gold < cost {
			fmt.Println("Pas assez d'or")
			continue
		}
		p.Gold = p.Gold - cost
		p = bag.AddBag(p, name)
		fmt.Println("Achat:", name, "Or restant:", p.Gold)
	}
}
