package main

import (
	"RED/Craft"
	"RED/Economie"
	hero "RED/Personnages"
	"RED/TourparTour"
	"fmt"
	"strings"
)

type Item struct {
	Name   string
	Type   string
	Effect string
	Slot   string
}

var Inventaire []Item

const CapaciteInventaire = 4

func main() {
	Inventaire = []Item{
		{Name: "Potion", Type: "consommable", Effect: "Restaure 20 PV", Slot: ""},
		{Name: "Épée rouillée", Type: "équipement", Effect: "+2 ATK", Slot: ""},
		{Name: "cuire"},
		{Name: "fils"},
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
			InfoPerso()
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
			fmt.Println("2 - Utiliser une potion")
			fmt.Println("3 - Passer le tour")
			fmt.Print("Choix : ")
			var choix int
			fmt.Scanln(&choix)

			switch choix {
			case 1:
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

			case 2:
				potionIndex := -1
				for j, item := range Inventaire {
					if item.Type == "consommable" && item.Name == "Potion" {
						potionIndex = j
						break
					}
				}

				if potionIndex != -1 {
					heal := 20
					team[i].PV += heal
					if team[i].PV > team[i].PVMax {
						team[i].PV = team[i].PVMax
					}
					fmt.Printf("%s utilise une potion et récupère %d PV !\n\n", team[i].Name, heal)
					Inventaire = append(Inventaire[:potionIndex], Inventaire[potionIndex+1:]...)
				} else {
					fmt.Println("Aucune potion disponible dans l'inventaire.\n")
				}

			case 3:
				fmt.Printf("%s passe son tour.\n\n", team[i].Name)

			default:
				fmt.Println("Choix invalide, tour perdu.\n")
			}
		}

		// Attaque du gobelin sur un héros vivant aléatoire
		TourparTour.GoblinPattern(&goblin, toHeroPointers(team), round)

		round++
	}

	if goblin.PV <= 0 {
		fmt.Println("Victoire des héros !")
	} else {
		fmt.Println("Le gobelin a gagné...")
	}
}

func PossedeIngredientsDansInventaire(ing1, ing2 string) bool {
	trouve1 := false
	trouve2 := false
	for _, item := range Inventaire {
		if item.Name == ing1 {
			trouve1 = true
		}
		if item.Name == ing2 {
			trouve2 = true
		}
	}
	return trouve1 && trouve2
}

func SupprimerItemInventaire(nom string) {
	for i, item := range Inventaire {
		if item.Name == nom {
			Inventaire = append(Inventaire[:i], Inventaire[i+1:]...)
			break
		}
	}
}

func InventairePlein() bool {
	return len(Inventaire) >= CapaciteInventaire
}

func FonctionForgeron() {
	fmt.Println("=== Forgeron ===")
	fmt.Println("Objets à fabriquer :")
	for i, item := range Craft.CraftItems {
		fmt.Printf("%d - %s (Recette: %s + %s)\n", i+1, item.Name, item.Name2, item.Name3)
	}

	var choix int
	fmt.Print("Quel objet veux-tu fabriquer ? ")
	fmt.Scanln(&choix)

	if choix >= 1 && choix <= len(Craft.CraftItems) {
		item := Craft.CraftItems[choix-1]

		if PossedeIngredientsDansInventaire(item.Name2, item.Name3) {
			SupprimerItemInventaire(item.Name2)
			SupprimerItemInventaire(item.Name3)

			Inventaire = append(Inventaire, Item{
				Name:   item.Name,
				Type:   item.Type,
				Effect: item.Effect,
				Slot:   item.Slot,
			})

			fmt.Println("✅", item.Name, "fabriqué et ajouté à l'inventaire.")
		} else {
			fmt.Println("❌ Tu n'as pas les bons ingrédients dans ton inventaire.")
		}
	} else {
		fmt.Println("Choix invalide.")
	}

	fmt.Println()
}

func FonctionSecondaire() {
	var choix int
	fmt.Println("=== SOUS-MENU ===")
	fmt.Println("1 - Marché")
	fmt.Println("2 - Forgeron")
	fmt.Println("0 - Retour au menu principal")
	fmt.Print("Ton choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
		fmt.Println("=== Marché ===")
		fmt.Printf("💰 Pièces : %d | 📦 Inventaire : %d/%d\n", Economie.Argent(), len(Inventaire), CapaciteInventaire)

		for i, item := range Economie.Market {
			fmt.Printf("%d - %s (Prix: %d pièces)\n", i+1, item.Name, item.Price)
		}

		offset := len(Economie.Market)
		if Economie.Market2Unlocked > 0 {
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
		} else if Economie.Market2Unlocked > 0 &&
			choix2 > len(Economie.Market) &&
			choix2 <= len(Economie.Market)+Economie.Market2Unlocked {
			item = Economie.Market2[choix2-len(Economie.Market)-1]
		} else {
			fmt.Println("❌ Numéro invalide.")
			return
		}

		resultat := Economie.Buy(item.Name)
		fmt.Println(resultat)

		if strings.HasPrefix(resultat, "Achat réussi") {
			if item.Type == "amélioration" && item.Name == "Extension d'inventaire" {
				CapaciteInventaire++
				fmt.Println("🧰 Capacité d'inventaire augmentée à", CapaciteInventaire)
				return
			}

			if InventairePlein() {
				fmt.Println("❌ Inventaire plein. Impossible d'ajouter l'objet.")
				return
			}

			Inventaire = append(Inventaire, Item{
				Name:   item.Name,
				Type:   item.Type,
				Effect: item.Effect,
				Slot:   item.Slot,
			})
		}

	case 2:
		fmt.Println("=== Forgeron ===")
		fmt.Printf("📦 Inventaire : %d/%d\n", len(Inventaire), CapaciteInventaire)
		fmt.Println("Objets à fabriquer :")
		for i, item := range Craft.CraftItems {
			fmt.Printf("%d - %s (Recette: %s + %s)\n", i+1, item.Name, item.Name2, item.Name3)
		}

		var choixForge int
		fmt.Print("Quel objet veux-tu fabriquer ? ")
		fmt.Scanln(&choixForge)

		if choixForge >= 1 && choixForge <= len(Craft.CraftItems) {
			item := Craft.CraftItems[choixForge-1]

			if InventairePlein() {
				fmt.Println("❌ Inventaire plein. Impossible de fabriquer l'objet.")
				return
			}

			if PossedeIngredientsDansInventaire(item.Name2, item.Name3) {
				SupprimerItemInventaire(item.Name2)
				SupprimerItemInventaire(item.Name3)

				Inventaire = append(Inventaire, Item{
					Name:   item.Name,
					Type:   item.Type,
					Effect: item.Effect,
					Slot:   item.Slot,
				})

				fmt.Println("✅", item.Name, "fabriqué et ajouté à l'inventaire.")
			} else {
				fmt.Println("❌ Tu n'as pas les bons ingrédients dans ton inventaire.")
			}
		} else {
			fmt.Println("❌ Choix invalide.")
		}

	case 0:
		fmt.Println("Retour au menu principal.")

	default:
		fmt.Println("❌ Choix invalide.")
	}

	fmt.Println()
}

func InfoPerso() {
	elise := hero.InitElise()
	jules := hero.InitJules()
	vittorio := hero.InitVittorio()

	heroes := []hero.Hero{*elise, *jules, *vittorio}

	fmt.Println("=== Informations des héros ===")
	for _, h := range heroes {
		fmt.Printf("Nom       : %s\n", h.Name)
		fmt.Printf("Classe    : %s\n", h.Classe)
		fmt.Printf("PV        : %d/%d\n", h.PV, h.PVMax)
		fmt.Printf("ATK       : %d\n", h.Atk)
		fmt.Printf("DEF       : %d\n", h.Def)
		fmt.Printf("Inventaire: %v\n", h.Inventory)
		fmt.Println("---------------------------")
	}
	fmt.Println()
}

func toHeroPointers(heroes []hero.Hero) []*hero.Hero {
	var ptrs []*hero.Hero
	for i := range heroes {
		ptrs = append(ptrs, &heroes[i])
	}
	return ptrs
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

	fmt.Println("0 - Retour")
	fmt.Print("Ton choix : ")
	var choix int
	fmt.Scanln(&choix)

	if choix != 1 {
		return
	}

	// Vérifier la présence de potion
	potionIndex := -1
	for i, item := range Inventaire {
		if item.Type == "consommable" && item.Name == "Potion" {
			potionIndex = i
			break
		}
	}

	if potionIndex == -1 {
		fmt.Println("❌ Aucune potion disponible dans l'inventaire.")
		return
	}

	// Afficher les héros disponibles
	heroes := []hero.Hero{
		*hero.InitElise(),
		*hero.InitJules(),
		*hero.InitVittorio(),
	}

	fmt.Println("Choisis un héros à soigner :")
	for i, h := range heroes {
		fmt.Printf("%d - %s (PV: %d/%d)\n", i+1, h.Name, h.PV, h.PVMax)
	}

	fmt.Print("Numéro du héros : ")
	fmt.Scanln(&choix)

	if choix >= 1 && choix <= len(heroes) {
		heal := 20
		heroes[choix-1].PV += heal
		if heroes[choix-1].PV > heroes[choix-1].PVMax {
			heroes[choix-1].PV = heroes[choix-1].PVMax
		}
		fmt.Printf("✅ %s récupère %d PV !\n", heroes[choix-1].Name, heal)
		Inventaire = append(Inventaire[:potionIndex], Inventaire[potionIndex+1:]...)
	} else {
		fmt.Println("Choix invalide.")
	}

	fmt.Println()
}
