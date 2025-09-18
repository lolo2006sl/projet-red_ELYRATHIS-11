package main

import (
	"RED/Craft"
	"RED/Economie"
	hero "RED/Personnages"
	"RED/TourparTour"
	"fmt"
)

type Item struct {
	Name   string
	Type   string
	Effect string
	Slot   string
}

var Inventaire []Item

func main() {
	Inventaire = []Item{
		{Name: "Potion", Type: "consommable", Effect: "Restaure 20 PV", Slot: ""},
		{Name: "Épée rouillée", Type: "équipement", Effect: "+2 ATK", Slot: ""},
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
		fmt.Println()

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
	team := []hero.Hero{
		*hero.InitElise(),
		*hero.InitJules(),
		*hero.InitVittorio(),
	}

	goblin := TourparTour.InitGoblin()
	round := 1

	for goblin.PV > 0 && TourparTour.AnyHeroAlive(team) {
		fmt.Println("=== Tour", round, "===")
		for _, h := range team {
			status := ""
			if h.PV <= 0 {
				status = " ⚠️ À terre"
			}
			fmt.Printf("%s - PV: %d/%d%s\n", h.Name, h.PV, h.PVMax, status)
		}
		fmt.Printf("Gobelin - PV: %d/%d\n\n", goblin.PV, goblin.PVMax)

		for i := range team {
			if team[i].PV <= 0 {
				continue
			}
			fmt.Printf("Tour de %s\n", team[i].Name)
			fmt.Println("1 - Attaquer")
			fmt.Println("2 - Passer le tour")
			fmt.Print("Choix : ")
			var choix int
			fmt.Scanln(&choix)

			if choix == 1 {
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
			} else if choix == 2 {
				fmt.Printf("%s passe son tour.\n\n", team[i].Name)
			} else {
				fmt.Println("Choix invalide, tour perdu.")
			}
		}

		for i := range team {
			if team[i].PV > 0 {
				oldPV := team[i].PV
				TourparTour.GoblinPattern(&goblin, &team[i], round)
				if team[i].PV <= 0 && oldPV > 0 {
					fmt.Printf("💀 %s est à terre !\n", team[i].Name)
				}
				break
			}
		}

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

	if choix == 1 {
		fmt.Println("=== Marché ===")
		fmt.Println("Vous avez :", Economie.Argent(), "pièces")

		for i, item := range Economie.Market {
			fmt.Printf("%d - %s (Prix: %d pièces)\n", i+1, item.Name, item.Price)
		}

		offset := len(Economie.Market)
		if Economie.Market2Unlocked == len(Economie.Market2) {
			fmt.Println("Objets spéciaux débloqués :")
			for i := 0; i < Economie.Market2Unlocked; i++ {
				item := Economie.Market2[i]
				fmt.Printf("%d - %s (Prix: %d pièces)\n", offset+i+1, item.Name, item.Price)
			}
		}

		var choix2 int
		fmt.Print("Entrez le numéro de l'item que vous voulez acheter : ")
		fmt.Scanln(&choix2)

		var item Economie.Item_market
		if choix2 >= 1 && choix2 <= len(Economie.Market) {
			item = Economie.Market[choix2-1]
		} else if Economie.Market2Unlocked == len(Economie.Market2) &&
			choix2 > len(Economie.Market) &&
			choix2 <= len(Economie.Market)+Economie.Market2Unlocked {
			item = Economie.Market2[choix2-len(Economie.Market)-1]
		} else {
			fmt.Println("Numéro invalide.")
			return
		}

		resultat := Economie.Buy(item.Name)
		fmt.Println(resultat)

		if len(resultat) >= 13 && resultat[:13] == "Achat réussi" {
			Inventaire = append(Inventaire, Item{
				Name:   item.Name,
				Type:   item.Type,
				Effect: item.Effect,
				Slot:   item.Slot,
			})
		}
	}

	if choix == 2 {
		fmt.Println("=== Craft ===")
		for _, item := range Craft.CraftItems {
			fmt.Println("Item :", item.Name, "| Recette :", item.Name2, "+", item.Name3)
		}
	}

	if choix == 0 {
		fmt.Println("Retour au menu principal.")
	}

	if choix != 0 && choix != 1 && choix != 2 {
		fmt.Println("Choix invalide.")
	}

	fmt.Println()
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
	fmt.Println()
}
