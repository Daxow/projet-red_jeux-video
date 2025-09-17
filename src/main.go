package main

import (
	"fmt"
	"projetred/src/player"
	"projetred/src/ui"
)

func main() {
	fmt.Println("Bienvenue")
	p := player.CreateCharacter()
	p = ui.Menu(p)
	fmt.Println("Fin du jeu")
}
