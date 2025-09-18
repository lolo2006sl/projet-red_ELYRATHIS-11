package TourparTour

import (
	hero "RED/Personnages"
	"fmt"
	"math/rand"
	"time"
)

func GoblinPattern(goblin *Monster, heroes []*hero.Hero, round int) {
	rand.Seed(time.Now().UnixNano())

	var aliveHeroes []*hero.Hero
	for _, h := range heroes {
		if h.PV > 0 {
			aliveHeroes = append(aliveHeroes, h)
		}
	}

	if len(aliveHeroes) == 0 {
		fmt.Println("Tous les héros sont morts.")
		return
	}

	cible := aliveHeroes[rand.Intn(len(aliveHeroes))]

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
	if cible.PV == 0 {
		cible.Wasted = true
	}

}
