package TourparTour

import (
	hero "RED/Persdonnages"
	"fmt"
)

// Fonction principale du combat tour par tour
func LancerCombat(joueur hero.Hero, ennemi Monster) {
	var choix int
	round := 1

	for joueur.PV > 0 && ennemi.PV > 0 {
		fmt.Println()
		fmt.Println("Tour", round)
		fmt.Printf("PV %s : %d/%d | PV %s : %d\n", joueur.Name, joueur.PV, joueur.PVMax, ennemi.Name, ennemi.PV)

		for {
			fmt.Println("Tape 1 pour attaquer")
			fmt.Scanln(&choix)
			if choix == 1 {
				break
			}
		}

		// Attaque du héros
		damageToMonster := joueur.Atk - ennemi.Def
		if damageToMonster < 0 {
			damageToMonster = 0
		}
		ennemi.PV -= damageToMonster
		if ennemi.PV < 0 {
			ennemi.PV = 0
		}
		fmt.Printf("%s attaque %s et inflige %d dégâts.\n", joueur.Name, ennemi.Name, damageToMonster)

		// Attaque du monstre
		GoblinPattern(&ennemi, &joueur, round)

		round++
	}

	fmt.Println()
	if joueur.PV > 0 {
		fmt.Println("Victoire du héros !")
	} else {
		fmt.Println("Le gobelin a gagné...")
	}
}
