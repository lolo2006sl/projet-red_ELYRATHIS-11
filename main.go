package main

import (
	"RED/Craft"
	"RED/Economie"
	hero "RED/Personnages"
	"RED/TourparTour"
	"fmt"
	"strconv"
	"strings"
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
var Elise = hero.InitElise()
var Jules = hero.InitJules()
var Vittorio = hero.InitVittorio()

var RecompenseVictoire int = 50 // ou une autre valeur selon ton système économique

func main() {
	fmt.Println(`
    ____  ____  __  ____  _________   ____  ______   __  ______    ____  _____ ____________    __    ______
   / __ )/ __ \/ / / /  |/  / ____/  / __ \/ ____/  /  |/  /   |  / __ \/ ___// ____/  _/ /   / /   / ____/
  / __  / /_/ / / / / /|_/ / __/    / / / / __/    / /|_/ / /| | / /_/ /\__ \/ __/  / // /   / /   / __/   
 / /_/ / _, _/ /_/ / /  / / /___   / /_/ / /___   / /  / / ___ |/ _, _/___/ / /____/ // /___/ /___/ /___   
/_____/_/ |_|\____/_/  /_/_____/  /_____/_____/  /_/  /_/_/  |_/_/ |_|/____/_____/___/_____/_____/_____/   
                                                                                                           `)
	fmt.Println("")

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
	team := []*hero.Hero{Elise, Jules, Vittorio}
	goblin := TourparTour.InitGoblin()
	round := 1

	// Réanimation des héros KO
	for i := range team {
		if team[i].Wasted {
			hero.ResetPV(team[i])
			team[i].Wasted = false
			fmt.Printf("%s revient avec %d PV.\n", team[i].Name, team[i].PV)
		}
	}

	for goblin.PV > 0 && TourparTour.AnyHeroAlivePtrs(team) {
		fmt.Println("=== Tour", round, "===")
		fmt.Println("PV du Gobelin :", goblin.PV)
		for _, h := range team {
			if h.PV > 0 {
				fmt.Printf("PV de %s : %d/%d\n", h.Name, h.PV, h.PVMax)
			} else {
				fmt.Printf("%s est à terre.\n", h.Name)
			}
		}

		for i := range team {
			if team[i].PV <= 0 {
				fmt.Println(team[i].Name, "est à terre et passe son tour.")
				continue
			}
			fmt.Println(team[i].Name, "entre en action.")
			fmt.Println("1 - Attaquer")
			fmt.Println("2 - Utiliser un skill")
			fmt.Println("3 - Utiliser une potion")
			fmt.Println("4 - Ne rien faire")
			fmt.Print("choix : ")
			var choix int
			fmt.Scanln(&choix)

			switch choix {
			case 1:
				degats := team[i].Atk
				goblin.PV -= degats
				if goblin.PV < 0 {
					goblin.PV = 0
				}
				fmt.Println()
				fmt.Println(team[i].Name, "attaque et inflige", degats, "dégâts au gobelin.")

			case 2:
				if len(team[i].Skill) == 0 {
					fmt.Println("Ce héros ne connaît aucun skill.")
					break
				}
				fmt.Println("Skills disponibles :")
				for j, s := range team[i].Skill {
					fmt.Printf("%d - %s\n", j+1, s)
				}
				var choixSkill int
				fmt.Print("Choisis un skill : ")
				fmt.Scanln(&choixSkill)
				if choixSkill >= 1 && choixSkill <= len(team[i].Skill) {
					skill := team[i].Skill[choixSkill-1]
					TourparTour.UtiliserSkill(team[i], skill, &goblin)
				} else {
					fmt.Println("Choix de skill invalide.")
				}

			case 3:
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
					fmt.Printf("%s utilise une potion et récupère %d PV !\n", team[i].Name, heal)
					Inventaire = append(Inventaire[:potionIndex], Inventaire[potionIndex+1:]...)
				} else {
					fmt.Println("Aucune potion disponible dans l'inventaire.")
				}

			case 4:
				fmt.Println(team[i].Name, "reste en retrait.")

			default:
				fmt.Println("Choix invalide. Le héros perd son tour.")
			}
		}

		// Le gobelin attaque
		cible := TourparTour.ChoisirCiblePtrs(team)
		if cible != nil {
			degats := goblin.Atk
			cible.PV -= degats
			fmt.Println("Le gobelin attaque", cible.Name, "et inflige", degats, "dégâts.")
		}
		round++
	}

	if goblin.PV <= 0 {
		fmt.Println("Victoire des héros !")
		Economie.AddMoney(RecompenseVictoire)
		fmt.Printf("Vous gagnez %d pièces ! Vous avez maintenant %d pièces.\n", RecompenseVictoire, Economie.Argent())
	} else {
		fmt.Println("Le gobelin a gagné...")
	}

	// Marquer les héros KO
	for i := range team {
		if team[i].PV <= 0 {
			team[i].Wasted = true
			fmt.Printf("%s est KO et sera réanimé au prochain combat.\n", team[i].Name)
		} else {
			team[i].Wasted = false
		}
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

	switch {
	case choix < 1 || choix > len(Craft.CraftItems):
		fmt.Println("❌ Choix invalide.")
		return

	default:
		objet := Craft.CraftItems[choix-1]

		switch PossedeIngredients(objet.Name2, objet.Name3) {
		case true:
			SupprimerIngredient(objet.Name2)
			SupprimerIngredient(objet.Name3)

			Inventaire = append(Inventaire, Item{
				Name:   objet.Name,
				Type:   objet.Type,
				Effect: objet.Effect,
				Slot:   objet.Slot,
			})

			fmt.Println("✅", objet.Name, "fabriqué et ajouté à l'inventaire.")

		case false:
			fmt.Println("❌ Tu n'as pas les bons ingrédients.")
		}
	}

	fmt.Println()
}

func SupprimerIngredient(nom string) {
	for i, item := range Inventaire {
		if item.Name == nom {
			Inventaire = append(Inventaire[:i], Inventaire[i+1:]...)
			break
		}
	}
}

func PossedeIngredients(nom1, nom2 string) bool {
	compteur := map[string]int{}

	for _, item := range Inventaire {
		compteur[item.Name]++
	}

	switch {
	case nom1 == nom2:
		return compteur[nom1] >= 2
	default:
		return compteur[nom1] >= 1 && compteur[nom2] >= 1
	}
}

func FonctionSecondaire() {
	var choix int
	fmt.Println("=== SOUS-MENU ===")
	fmt.Println("1 - Market")
	fmt.Println("2 - Forgeron")
	fmt.Println("0 - Retour au menu principal")
	fmt.Print("Ton choix : ")
	fmt.Scanln(&choix)

	switch choix {
	case 1:
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

		totalMarket := len(Economie.Market)
		totalMarket2 := len(Economie.Market2)

		if choix2 >= 1 && choix2 <= totalMarket {
			item = Economie.Market[choix2-1]
		} else if Economie.Market2Unlocked > 0 &&
			choix2 > totalMarket &&
			choix2 <= totalMarket+Economie.Market2Unlocked &&
			Economie.Market2Unlocked <= totalMarket2 {
			item = Economie.Market2[choix2-totalMarket-1]
		} else {
			fmt.Println("Numéro invalide ou marché secondaire inaccessible.")
			return
		}

		resultat := Economie.Buy(item.Name)
		fmt.Println(resultat)

		if len(resultat) >= 13 && resultat[:13] == "Achat réussi" {
			if item.Type == "amélioration" && item.Name == "Extension d'inventaire" {

				fmt.Println("Capacité d'inventaire augmentée à", CapaciteInventaire)
				return
			}

			if InventairePlein() {
				fmt.Println("Inventaire plein. Impossible d'ajouter l'objet.")
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
		FonctionForgeron()
	default:
		fmt.Println("Choix invalide")
	}
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
		if item.Type == "équipement" {
			fmt.Println("À quel héros veux-tu équiper cet objet ?")
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

			// Appliquer l'effet "+ 10 PV"
			effet := strings.TrimSpace(item.Effect)
			if strings.HasPrefix(effet, "+") && strings.Contains(effet, "PV") {
				parts := strings.Split(effet, " ")
				if len(parts) >= 2 {
					val, err := strconv.Atoi(strings.TrimSpace(parts[1]))
					if err == nil {
						cible.PVMax += val
						cible.PV += val
						if cible.PV > cible.PVMax {
							cible.PV = cible.PVMax
						}
						fmt.Printf("✅ %s a équipé %s. PV Max : %d | PV actuel : %d\n", cible.Name, item.Name, cible.PVMax, cible.PV)
					}
				}
			} else {
				fmt.Printf("✅ %s a équipé %s. Effet : %s\n", cible.Name, item.Name, item.Effect)
			}

			SupprimerItemInventaire(item.Name)
			return
		}
	}
}
