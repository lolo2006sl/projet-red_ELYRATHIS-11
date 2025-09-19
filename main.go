package main

import (
	"RED/Craft"
	"RED/Economie"
	hero "RED/Personnages"
	"RED/TourparTour"
	"fmt"
	"time"
)

type Item struct {
	Name   string
	Type   string
	Effect string
	Slot   string
}

const CapaciteInventaire = 10

var Inventaire []Item
var Elise *hero.Hero = hero.InitElise()
var Jules *hero.Hero = hero.InitJules()
var Vittorio *hero.Hero = hero.InitVittorio()

func main() {
	Inventaire = []Item{
		{Name: "Potion", Type: "consommable", Effect: "Restaure 20 PV", Slot: ""},
		{Name: "Épée rouillée", Type: "équipement", Effect: "+2 ATK", Slot: ""},
		{Name: "cuire", Type: "ressource", Effect: "", Slot: ""},
		{Name: "fils", Type: "ressource", Effect: "", Slot: ""},
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
		*Elise,
		*Jules,
		*Vittorio,
	}

	goblin := TourparTour.InitGoblin()
	round := 1

	for goblin.PV > 0 && TourparTour.AnyHeroAlive(team) {
		fmt.Println("=== Tour", round, "===")

		// ... tout le reste du combat ...

		round++
	}

	if goblin.PV <= 0 {
		fmt.Println("Victoire des héros !")
	} else {
		fmt.Println("Le gobelin a gagné...")
	}
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
	fmt.Println("1 - Market")
	fmt.Println("2 - Forgeron")
	fmt.Println("0 - Retour au menu principal")
	fmt.Print("Ton choix : ")
	fmt.Scanln(&choix)

	if choix == 1 {
		fmt.Println("=== Marché ===")
		fmt.Printf("Vous avez : %d pièces | Inventaire : %d/%d\n", Economie.Argent(), len(Inventaire), CapaciteInventaire)

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
			fmt.Println("Numéro invalide.")
			return
		}

		resultat := Economie.Buy(item.Name)
		fmt.Println(resultat)

		if len(resultat) >= 13 && resultat[:13] == "Achat réussi" {
			if item.Type == "amélioration" && item.Name == "Extension d'inventaire" {

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
	}

	// ... (le reste de ta fonction continue ici, pour le forgeron et les autres choix)
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

	if choix == 0 || choix < 1 || choix > len(Inventaire) {
		return
	}

	item := Inventaire[choix-1]

	// 📘 Si c'est un livre de sort
	if item.Name == "Livre de Sort : Boule de Feu" {
		fmt.Println("À quel héros veux-tu enseigner le sort ?")
		fmt.Println("1 -", Elise.Name)
		fmt.Println("2 -", Jules.Name)
		fmt.Println("3 -", Vittorio.Name)
		var choixHero int
		fmt.Print("Numéro du héros : ")
		fmt.Scanln(&choixHero)

		switch choixHero {
		case 1:
			hero.SpellBook(Elise)
			fmt.Println("📘", Elise.Name, "a appris le sort 'Boule de Feu' !")
		case 2:
			hero.SpellBook(Jules)
			fmt.Println("📘", Jules.Name, "a appris le sort 'Boule de Feu' !")
		case 3:
			hero.SpellBook(Vittorio)
			fmt.Println("📘", Vittorio.Name, "a appris le sort 'Boule de Feu' !")
		default:
			fmt.Println("Choix invalide.")
			return
		}

		SupprimerItemInventaire(item.Name)
		return
	}

	// 🧪 Si c'est une potion
	if item.Type == "consommable" {
		fmt.Println("Choisis un héros cible :")
		fmt.Println("1 -", Elise.Name)
		fmt.Println("2 -", Jules.Name)
		fmt.Println("3 -", Vittorio.Name)
		var choixHero int
		fmt.Print("Numéro du héros : ")
		fmt.Scanln(&choixHero)

		var cible *hero.Hero
		switch choixHero {
		case 1:
			cible = Elise
		case 2:
			cible = Jules
		case 3:
			cible = Vittorio
		default:
			fmt.Println("Choix invalide.")
			return
		}

		if item.Effect == "Restaure 20 PV" {
			heal := 20
			cible.PV += heal
			if cible.PV > cible.PVMax {
				cible.PV = cible.PVMax
			}
			fmt.Printf("🧪 %s utilise %s et récupère %d PV. PV : %d/%d\n", cible.Name, item.Name, heal, cible.PV, cible.PVMax)
			SupprimerItemInventaire(item.Name)
			return
		} else if item.Name == "Potion de poison" {
			fmt.Printf("🧪 %s utilise %s... mais quelque chose cloche.\n", cible.Name, item.Name)
			go func() {
				for i := 0; i < 5; i++ {
					time.Sleep(3 * time.Second)
					cible.PV -= 5
					if cible.PV < 0 {
						cible.PV = 0
					}
					fmt.Printf("☠️ %s subit 5 dégâts de poison. PV restant : %d\n", cible.Name, cible.PV)
				}
			}()
			SupprimerItemInventaire(item.Name)
			return
		} else {
			fmt.Println("❌ Effet inconnu pour cet objet.")
		}
	}
}
