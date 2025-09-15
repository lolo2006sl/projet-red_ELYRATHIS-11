package main

import (
	"RED/TourparTour"
	"fmt"
)

func main() {
	hero := TourparTour.InitFakeHero()
	goblin := TourparTour.InitGoblin()

	var choix int
	round := 1

	for hero.PV > 0 && goblin.PV > 0 {
		fmt.Println()
		fmt.Println("Tour", round)
		// Affichage des PV
		fmt.Println("PV", hero.Name, ":", hero.PV, "| PV", goblin.Name, ":", goblin.PV)
		// choix du joueur
		for {
			fmt.Println("Tape 1 pour attaquer")
			fmt.Scanln(&choix)

			if choix == 1 { // check si le joueur a tapé 1
				break
			}
		}

		// Héros attaque
		damageToGoblin := hero.CalculateDamage(goblin.Def)
		goblin.PV -= damageToGoblin
		if goblin.PV < 0 {
			goblin.PV = 0
		}
		fmt.Println(hero.Name, "attaque", goblin.Name, "et inflige", damageToGoblin, "dégâts") //print les dégats

		// Gobelin attaque
		damageToHero := goblin.CalculateDamage(hero.Def)
		hero.PV -= damageToHero
		if hero.PV < 0 {
			hero.PV = 0
		}
		fmt.Println(goblin.Name, "attaque", hero.Name, "et inflige", damageToHero, "dégâts") //print les dégats

		round++
	}

	fmt.Println()
	if hero.PV > 0 {
		fmt.Println("Victoire du héros !")
	} else {
		fmt.Println("Le gobelin a gagné...")

	}
}
