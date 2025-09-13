package main

import (
	"RED/TourparTour"
)

func main() {
	goblin := TourparTour.InitGoblin()

	goblin.DisplayHP()

	damage := goblin.CalculateDamage()
	println("Dégâts infligés par le gobelin :", damage)

	goblin.DisplayHP()
}
