package TourparTour

import (
	hero "RED/Personnages"
	"fmt"
)

func GoblinPattern(g *Monster, target *hero.Hero, turn int) {
	var damage int
	if turn%3 == 0 {
		damage = g.Atk * 2
		fmt.Printf("%s utilise une attaque spéciale sur %s et inflige %d dégâts !\n", g.Name, target.Name, damage)
	} else {
		damage = g.CalculateDamage(target.Def)
		fmt.Printf("%s attaque %s et inflige %d dégâts.\n", g.Name, target.Name, damage)
	}
	target.PV -= damage
	if target.PV <= 0 {
		target.PV = 1
	}
}
