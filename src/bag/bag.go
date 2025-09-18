package bag

import (
	"fmt"
	"projetred/src/data"
	"projetred/src/player"
	"time"
)

var m data.Mob

func ShowBag(p data.Person) {
	fmt.Println("\nSac:")
	fmt.Print("\n")
	for i := 0; i < len(p.Bag); i++ {
		fmt.Printf("%d. %s\n", i+1, p.Bag[i])
	}
}

func AddBag(p data.Person, it string) data.Person {
	if len(p.Bag) >= p.Size {
		fmt.Println("\nSac plein")
		return p
	}
	p.Bag = append(p.Bag, it)
	return p
}

func RemoveBag(p data.Person, pos int) data.Person {
	if pos < 0 || pos >= len(p.Bag) {
		return p
	}
	n := []string{}
	for i := 0; i < len(p.Bag); i++ {
		if i != pos {
			n = append(n, p.Bag[i])
		}
	}
	p.Bag = n
	return p
}

func UseBag(p data.Person, m data.Mob) data.Person {
	if len(p.Bag) == 0 {
		fmt.Println("\nSac vide")
		return p
	}
	ShowBag(p)
	fmt.Println("\nVotre choix?")
	var x int
	fmt.Scanln(&x)
	if x <= 0 || x > len(p.Bag) {
		return p
	}

	it := p.Bag[x-1]

	if it == "Potion de vie" {
		p.Hp = p.Hp + 50
		if p.Hp > p.Hpmax {
			p.Hp = p.Hpmax
		}
		fmt.Println("\nVous utilisez Potion de vie")
		fmt.Println("PV:", p.Hp, "/", p.Hpmax)
		p = RemoveBag(p, x-1)
		return p
	}

	if it == "Potion de poison" {
		fmt.Println("\nVous utilisez Potion de poison")
		for i := 0; i < 3; i++ {
			time.Sleep(1 * time.Second)
			m.Hp -= 10
			if m.Hp < 0 {
				m.Hp = 0
			}
			fmt.Println("Poison:", m.Hp, "/", m.Hpmax)
			if m.Hp <= 0 {
				fmt.Println("Le monstre est mort !")
				break
			}
		}
		p = RemoveBag(p, x-1)
		return p
	}

	if it == "Augmentation d’inventaire" {
		if p.Up >= 3 {
			fmt.Println("Plus d’augmentation possible")
			return p
		}
		p.Size = p.Size + 10
		p.Up = p.Up + 1
		fmt.Println("Capacité du sac:", p.Size)
		p = RemoveBag(p, x-1)
		return p
	}

	if it == "Livre de Sort : Boule de Feu" {
		p = player.SpellBook(p)
		p = RemoveBag(p, x-1)
		return p
	}

	if it == "Chapeau de l’aventurier" || it == "Tunique de l’aventurier" || it == "Bottes de l’aventurier" {
		p = player.Equip(p, it)
		fmt.Println("Équipé:", it)
		p = RemoveBag(p, x-1)
		return p
	}

	fmt.Println("Impossible d'utiliser", it)
	return p
}
