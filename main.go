package main

import (
	"RED/Craft"
	"RED/Economie"
	hero "RED/Personnages"
	"RED/TourparTour"
	"fmt"
)

func main() {
	var MENU int

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1 - Lancer le combat")
		fmt.Println("2 - Autre fonction (économie/craft)")
		fmt.Println("3 - Afficher les héros et tester une attaque")
		fmt.Println("0 - Quitter")
		fmt.Print("Ton choix : ")
		fmt.Scanln(&MENU)

		if MENU == 1 {
			LancerCombat()
		} else if MENU == 2 {
			FonctionSecondaire()
		} else if MENU == 3 {
			TestAttaque()
		} else if MENU == 0 {
			fmt.Println("À bientôt !")
			return
		} else {
			fmt.Println("Choix invalide")
		}
	}
}

func LancerCombat() {
	hero := TourparTour.InitFakeHero()
	goblin := TourparTour.InitGoblin()

	var choix int
	round := 1

	for hero.PV > 0 && goblin.PV > 0 {
		fmt.Println("Tour", round)
		fmt.Printf("PV %s : %d / %d | PV %s : %d / %d", hero.Name, hero.PV, hero.PVMax, goblin.Name, goblin.PV, goblin.PVMax)

		for {
			fmt.Println("Tape 1 pour attaquer")
			fmt.Scanln(&choix)
			if choix == 1 {
				break
			}
		}

		// Héros attaque
		damageToGoblin := hero.CalculateDamage(goblin.Def)
		goblin.PV -= damageToGoblin
		if goblin.PV < 0 {
			goblin.PV = 0
		}
		fmt.Printf("%s attaque %s et inflige %d dégâts", hero.Name, goblin.Name, damageToGoblin)

		// Gobelin attaque avec GoblinPattern
		TourparTour.GoblinPattern(&goblin, &hero, round)

		round++
	}

	fmt.Println()
	if hero.PV > 0 {
		fmt.Println("Victoire du héros !")
	} else {
		fmt.Println("Le gobelin a gagné...")
	}
}

func FonctionSecondaire() {
	var choix int
	fmt.Println("=== SOUS-MENU ===")
	fmt.Println("1 - Market")
	fmt.Println("2 - Craft")
	fmt.Println("0 - Retour au menu principal")
	fmt.Print("Ton choix : ")
	fmt.Scanln(&choix)

	if choix == 1 {
		nom := Economie.Market[0].Name
		fmt.Println("Nom du premier item :", nom)
		price, found := Economie.GetPrice("Rubis")
		if found {
			fmt.Println("Prix du Rubis :", price)
		} else {
			fmt.Println("Item non trouvé")
		}
	} else if choix == 2 {
		for _, item := range Craft.CraftItems {
			fmt.Println("Item :", item.Name, "| Détail :", item.Name2)
		}
	} else if choix == 0 {
		fmt.Println("Retour au menu principal.")
	} else {
		fmt.Println("Choix invalide.")
	}
}

// perso
func TestAttaque() {
	elise := hero.InitElise()
	jules := hero.InitJules()
	vittorio := hero.InitVittorio()

	fmt.Println("Héros disponibles :")
	fmt.Printf("%s (%s) - PV: %d/%d", elise.Name, elise.Classe, elise.PV, elise.PVMax)
	fmt.Printf("%s (%s) - PV: %d/%d", jules.Name, jules.Classe, jules.PV, jules.PVMax)
	fmt.Printf("%s (%s) - PV: %d/%d", vittorio.Name, vittorio.Classe, vittorio.PV, vittorio.PVMax)

	Attaquer(jules, vittorio)
	fmt.Printf("Après l’attaque : %s a %d PV", vittorio.Name, vittorio.PV)
}

func Attaquer(attacker, target *hero.Hero) {
	degats := attacker.Atk - target.Def
	if degats < 0 {
		degats = 0
	}
	target.PV -= degats
	if target.PV < 0 {
		target.PV = 0
	}
	fmt.Printf("%s attaque %s et inflige %d dégâts !", attacker.Name, target.Name, degats)
}
