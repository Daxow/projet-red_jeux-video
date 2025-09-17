package combat

import (
	"fmt"
	"projetred/src/bag"
	"projetred/src/data"
	"projetred/src/player"
)

func MakeGoblin() data.Mob {
	m := data.Mob{}
	m.Name = "Gobelin"
	m.Hp = 40
	m.Hpmax = 40
	m.Atk = 5
	return m
}

func Fight(p data.Person) data.Person {
	m := MakeGoblin()
	fmt.Println(m.Hp)
	fmt.Println(m.Hpmax)
	t := 1
	for {
		fmt.Println("--- Tour", t, "---")
		fmt.Println("1. Attaque\n2. Sac\n0. Quittter")
		var x int
		fmt.Scanln(&x)

		if x == 0 {
			return p
		}
		if x == 1 {
			m.Hp = m.Hp - 5
			fmt.Println("Vous attaquez", m.Name, "il lui reste", m.Hp, "/", m.Hpmax)
		}
		if x == 2 {
			p = bag.UseBag(p, m)
		}
		if m.Hp <= 0 {
			fmt.Println("Victoire !")
			break
		}

		p.Hp = p.Hp - m.Atk
		fmt.Println(m.Name, "vous inflige", m.Atk, "il vous reste", p.Hp, "/", p.Hpmax)
		p = player.Dead(p)
		if p.Hp <= 0 {
			break
		}
		t = t + 1
	}
	return p
}
