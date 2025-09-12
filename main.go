package main

import (
	TPT "RED/TourparTour"
	"fmt"
)

func main() {
	goblin := TPT.InitGoblin()

	fmt.Println("=== Gobelin ===")
	fmt.Printf("Nom     : %s\n", goblin.Name)
	fmt.Printf("PV      : %d / %d\n", goblin.PV, goblin.PVMax)
	fmt.Printf("Défense : %d\n", goblin.Def)
	fmt.Printf("Attaque : %d\n", goblin.Atk)
}
