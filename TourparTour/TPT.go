package TourparTour

import (
	hero "RED/Personnages"
	"fmt"
	"math/rand"
	"time"
)

// Fonction principale du combat tour par tour (ancienne version solo)
func LancerCombat(joueur hero.Hero, ennemi Monster) {
	var choix int
	round := 1

	for joueur.PV > 0 && ennemi.PV > 0 {
		fmt.Println()
		fmt.Println("=== Tour", round, "===")

		status := ""
		if joueur.PV <= 0 {
			status = " - À terre"
		}
		fmt.Printf("PV %s : %d/%d%s | PV %s : %d/%d\n", joueur.Name, joueur.PV, joueur.PVMax, status, ennemi.Name, ennemi.PV, ennemi.PVMax)

		for {
			fmt.Println("Tape 1 pour attaquer")
			fmt.Scanln(&choix)
			if choix == 1 {
				break
			}
		}

		// Attaque du héros
		damageToMonster := joueur.Atk - ennemi.Def
		if damageToMonster <= 0 {
			damageToMonster = 1
		}
		ennemi.PV -= damageToMonster
		if ennemi.PV < 0 {
			ennemi.PV = 0
		}
		fmt.Printf("%s attaque %s et inflige %d dégâts.\n", joueur.Name, ennemi.Name, damageToMonster)

		// Attaque du monstre
		GoblinPattern(&ennemi, []*hero.Hero{&joueur}, round)

		if joueur.PV <= 0 {
			fmt.Printf("%s est à terre.\n", joueur.Name)
		}

		round++
	}

	fmt.Println()
	if joueur.PV > 0 {
		fmt.Println("Victoire du héros !")
	} else {
		fmt.Println("Le gobelin a gagné.")
	}
	fmt.Println("")
}

// Attaque du gobelin sur un héros vivant aléatoire
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
}

// Vérifie si au moins un héros est encore vivant
func AnyHeroAlive(team []hero.Hero) bool {
	for _, h := range team {
		if h.PV > 0 {
			return true
		}
	}
	return false
}
