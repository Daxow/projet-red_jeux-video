package player

import (
	"projetred/src/data"
	"fmt"
)

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

func UseSpell(spell string, p data.Person, m *data.Mob) {
	if spell == "Boule de feu" {
		m.Hp = m.Hp - 8
		fmt.Println("Vous lancez Boule de feu, il reste", m.Hp, "/", m.Hpmax)
	}
}