package player

import "projetred/src/data"

func Equip(p data.Person, item string) data.Person {
	if item == "Chapeau de l’aventurier" {
		p.Eq.Head = item
	}
	if item == "Tunique de l’aventurier" {
		p.Eq.Chest = item
	}
	if item == "Bottes de l’aventurier" {
		p.Eq.Feet = item
	}
	p = RecomputeHP(p)
	return p
}

func RecomputeHP(p data.Person) data.Person {
	base := 100
	if p.Class == "Elfe" {
		base = 80
	}
	if p.Class == "Nain" {
		base = 120
	}

	if p.Eq.Head == "Chapeau de l’aventurier" {
		base = base + 10
	}
	if p.Eq.Chest == "Tunique de l’aventurier" {
		base = base + 25
	}
	if p.Eq.Feet == "Bottes de l’aventurier" {
		base = base + 15
	}

	p.Hpmax = base
	if p.Hp > p.Hpmax {
		p.Hp = p.Hpmax
	}
	return p
}
