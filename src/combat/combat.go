package combat

import (
	"fmt"
	"projetred/src/bag"
	"projetred/src/data"
	"projetred/src/player"
)

func MakeGoblin() data.Mob {
	m := data.Mob{}
	m.Name = "Gobelin d’entrainement"
	m.Hp = 40
	m.Hpmax = 40
	m.Atk = 5
	return m
}

func Fight(p data.Person) data.Person {
	m := MakeGoblin()
	t := 1
	for {
		fmt.Println("--- Tour", t, "---")
		fmt.Println("1 Attaque 2 Sac 0 Quit")
		var x int
		fmt.Scanln(&x)

		if x == 0 {
			return p
		}
		if x == 1 {
			m.Hp = m.Hp - 5
			fmt.Println("Vous tapez", m.Name, "reste", m.Hp)
		}
		if x == 2 {
			p = bag.UseBag(p)
		}
		if m.Hp <= 0 {
			fmt.Println("Victoire !")
			break
		}

		p.Hp = p.Hp - m.Atk
		fmt.Println(m.Name, "tape", m.Atk, "reste", p.Hp)
		p = player.Dead(p)
		if p.Hp <= 0 {
			break
		}
		t = t + 1
	}
	return p
}
