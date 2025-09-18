package main

import (
	"RED/Craft"
	"RED/Economie"
	hero "RED/Personnages"
	"RED/TourparTour"
	"fmt"
)

// Structure de l'inventaire
type Item struct {
	Name   string
	Type   string
	Effect string
}

var Inventaire []Item

func main() {
	// Initialisation de l'inventaire
	Inventaire = []Item{
		{Name: "Potion", Type: "consommable", Effect: "Restaure 20 PV"},
		{Name: "Épée rouillée", Type: "équipement", Effect: "+2 ATK"},
	}

	var MENU int
	for {
		fmt.Println("=== MENU ===")
		fmt.Println("1 - Lancer le combat")
		fmt.Println("2 - économie et craft")
		fmt.Println("3 - info perso")
		fmt.Println("4 - inventaire")
		fmt.Println("0 - Quitter")
		fmt.Print("Ton choix : ")
		fmt.Scanln(&MENU)
		fmt.Println("")

		switch MENU {
		case 1:
			LancerCombat()
		case 2:
			FonctionSecondaire()
		case 3:
			TestAttaque()
		case 4:
			AfficherInventaire()
		case 0:
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func LancerCombat() {
	team := []hero.Personnage{
		hero.InitElise(),
		hero.InitJules(),
		hero.InitVittorio(),
	}

	goblin := TourparTour.InitGoblin()
	round := 1
	for goblin.PV > 0 && TourparTour.AnyHeroAlive(team) {
		fmt.Println("=== Tour", round, "===")
		for _, h := range team {
			fmt.Printf("%s - PV: %d/%d\n", h.Name, h.PV, h.PVMax)
		}
		fmt.Printf("Gobelin - PV: %d/%d\n\n", goblin.PV, goblin.PVMax)

		for i := range team {
			if team[i].PV <= 0 {
				continue
			}
			fmt.Printf("%s attaque %s\n", team[i].Name, goblin.Name)
			damage := team[i].Atk - goblin.Def
			if damage <= 0 {
				damage = 1
			}
			goblin.PV -= damage
			if goblin.PV < 0 {
				goblin.PV = 0
			}
			fmt.Printf("→ %s inflige %d dégâts\n\n", team[i].Name, damage)
		}

		TourparTour.GoblinPattern(&goblin, team, round)
		round++
	}

	if goblin.PV <= 0 {
		fmt.Println("Victoire des héros !")
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

	switch choix {
	case 1:
		for _, Market := range Economie.Market {
			fmt.Println("Item :", Market.Name, "| Recette:", Market.Price)
			fmt.Println("Vous avez: ", Economie.Argent())
		}
	case 2:
		for _, item := range Craft.CraftItems {
			fmt.Println("Item :", item.Name, "| Recette:", item.Name2, "| Recette:", item.Name3)
		}
	case 0:
		fmt.Println("Retour au menu principal.")
	default:
		fmt.Println("Choix invalide.")
	}
	fmt.Println("")
}

func TestAttaque() {
	elise := hero.InitElise()
	jules := hero.InitJules()
	vittorio := hero.InitVittorio()

	fmt.Println("Héros disponibles :")
	fmt.Printf("%s (%s) - PV: %d/%d\n", elise.Name, elise.Classe, elise.PV, elise.PVMax)
	fmt.Printf("%s (%s) - PV: %d/%d\n", jules.Name, jules.Classe, jules.PV, jules.PVMax)
	fmt.Printf("%s (%s) - PV: %d/%d\n\n", vittorio.Name, vittorio.Classe, vittorio.PV, vittorio.PVMax)
}

func AfficherInventaire() {
	if len(Inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	fmt.Println("=== Inventaire ===")
	for i, item := range Inventaire {
		fmt.Printf("%d - %s [%s] : %s\n", i+1, item.Name, item.Type, item.Effect)
	}
	fmt.Println("")
}
