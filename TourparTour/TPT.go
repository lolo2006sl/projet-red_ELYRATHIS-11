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

// Fonction pour utiliser un skill
func UtiliserSkill(h *hero.Hero, skill string, cible *Monster) {
	switch skill {
	case "Coup de poing":
		degats := h.Atk
		cible.PV -= degats
		fmt.Printf("%s utilise %s et inflige %d dégâts au gobelin.\n", h.Name, skill, degats)
	case "Boule de feu":
		degats := h.Atk*2 - cible.Def
		if degats < 0 {
			degats = 0
		}
		cible.PV -= degats
		fmt.Printf("%s lance %s et inflige %d dégâts magiques au gobelin !\n", h.Name, skill, degats)
	default:
		fmt.Println("Skill inconnu.")
	}
	if cible.PV < 0 {
		cible.PV = 0
	}
}

// ... autres fonctions comme LancerCombat, UtiliserSkill, etc.

func AnyHeroAlive(team []hero.Hero) bool {
	for _, h := range team {
		if h.PV > 0 {
			return true
		}
	}
	return false
}

// ✅ Ajoute cette fonction juste ici :
func ChoisirCible(team []hero.Hero) *hero.Hero {
	rand.Seed(time.Now().UnixNano())

	var aliveHeroes []*hero.Hero
	for i := range team {
		if team[i].PV > 0 {
			aliveHeroes = append(aliveHeroes, &team[i])
		}
	}

	if len(aliveHeroes) == 0 {
		return nil
	}

	return aliveHeroes[rand.Intn(len(aliveHeroes))]
}
