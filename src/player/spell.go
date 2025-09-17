package player

import "projetred/src/data"

func SpellBook(p data.Person) data.Person {
	has := false
	for i := 0; i < len(p.Spells); i++ {
		if p.Spells[i] == "Boule de feu" {
			has = true
		}
	}
	if !has {
		p.Spells = append(p.Spells, "Boule de feu")
	}
	return p
}
