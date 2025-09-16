package main

import (
    "RED/Economie"
    "RED/TourparTour"
    "RED/personnages"
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
            break
        } else if MENU == 2 {
            FonctionSecondaire()
        } else if MENU == 0 {
            fmt.Println("À bientôt !")
            return
        } else {
            fmt.Println("NON")
        }
    }

    joueur := hero.InitJules() // ou InitElise(), InitVittorio()
    ennemi := TourparTour.InitGoblin()

    TourparTour.LancerCombat(joueur, ennemi)
}

func FonctionSecondaire() {
    nom := Economie.Market[0].Name
    fmt.Println("Nom du premier item :", nom)
    price, found := Economie.GetPrice("Rubis")
    if found {
        fmt.Println("Prix du Rubis :", price)
    } else {
        fmt.Println("Item non trouvé")
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

    if choix {
        nom := Economie.Market[0].Name
        fmt.Println("Nom du premier item :", nom)
        price, found := Economie.GetPrice("Rubis")
        if found {
            fmt.Println("Prix du Rubis :", price)
        } else {
            fmt.Println("Item non trouvé")
        } else if 2: {
      		SystemCraft()
		}
		default:
        	fmt.Println("non Secondaire.")
	    }
	}
}