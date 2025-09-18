package forge

import (
	"fmt"
	"projetred/src/bag"
	"projetred/src/data"
	"projetred/src/player"
)

func hasItem(p data.Person, it string, used []int) (bool, int) {
	for i := 0; i < len(p.Bag); i++ {
		if p.Bag[i] == it {
			ok := true
			for j := 0; j < len(used); j++ {
				if used[j] == i {
					ok = false
				}
			}
			if ok {
				return true, i
			}
		}
	}
	return false, -1
}

func craft(p data.Person, need []string, item string) data.Person {
	if p.Gold < 5 {
		fmt.Println("Il faut 5 or")
		return p
	}
	used := []int{}
	ok := true
	for i := 0; i < len(need); i++ {
		found, pos := hasItem(p, need[i], used)
		if !found {
			ok = false
		}
		if found {
			used = append(used, pos)
		}
	}
	if !ok {
		fmt.Println("Ressources manquantes")
		return p
	}
	p.Gold = p.Gold - 5
	for i := 0; i < len(used); i++ {
		p = bag.RemoveBag(p, used[i]-i)
	}
	p = bag.AddBag(p, item)
	fmt.Println("Fabriqué:", item)
	return p
}

func Smith(p data.Person) data.Person {
	for {
		fmt.Println("\nForgeron (5 or):")
		fmt.Print("\n")
		fmt.Println("1. Chapeau de l’aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier)")
		fmt.Println("2. Tunique de l’aventurier (2 Fourrure de Loup, 1 Peau de Troll)")
		fmt.Println("3. Bottes de l’aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier)")
		fmt.Println("0. Retour")
		var x int
		fmt.Scanln(&x)
		if x == 0 {
			return p
		}

		if x == 1 {
			p = craft(p, []string{"Plume de Corbeau", "Cuir de Sanglier"}, "Chapeau de l’aventurier")
		}
		if x == 2 {
			p = craft(p, []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}, "Tunique de l’aventurier")
		}
		if x == 3 {
			p = craft(p, []string{"Fourrure de Loup", "Cuir de Sanglier"}, "Bottes de l’aventurier")
		}
		p = player.RecomputeHP(p)
	}
}
