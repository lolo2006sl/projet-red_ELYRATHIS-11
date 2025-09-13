package main

import (
	TPT "RED/TourparTour"
)

func main() {
	goblin := TPT.InitGoblin()

	goblin.DisplayHP()

	damage := goblin.CalculateDamage()
	println("Dégâts infligés par le gobelin :", damage)

	goblin.DisplayHP()
}
