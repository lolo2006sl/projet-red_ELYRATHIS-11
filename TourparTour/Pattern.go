package TourparTour

import (
	hero "RED/Personnages"
	"fmt"
)

func GoblinPattern(goblin *Monster, cible *hero.Hero, round int) {
	var damage int
	if round%3 == 0 {
		damage = goblin.Atk*2 - cible.Def
		fmt.Printf("%s utilise une attaque spéciale sur %s et inflige %d dégâts !\n", goblin.Name, cible.Name, damage)
	} else {
		damage = goblin.Atk - cible.Def
		fmt.Printf("%s attaque %s et inflige %d dégâts.\n", goblin.Name, cible.Name, damage)
	}

	if damage <= 0 {
		damage = 0
	}

	cible.PV -= damage
	if cible.PV < 0 {
		cible.PV = 0
	}
}
