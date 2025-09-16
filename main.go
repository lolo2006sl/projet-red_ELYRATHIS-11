package main

import (
	"RED/Craft"
	"RED/Economie"
	"RED/TourparTour"
	"fmt"
)

func main() {
	var MENU int

	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1 - Lancer le combat")
		fmt.Println("2 - Autre fonction (à venir)")
		fmt.Println("0 - Quitter")
		fmt.Print("Ton choix : ")
		fmt.Scanln(&MENU)

		if MENU == 1 {
			break // on sort du menu pour lancer le combat
		} else if MENU == 2 {
			FonctionSecondaire()
		} else if MENU == 0 {
			fmt.Println("À bientôt !")
			return
		} else {
			fmt.Println("NON") // message d'erreur si le choix est invalide
		}
	}

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

// fonction economie
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
