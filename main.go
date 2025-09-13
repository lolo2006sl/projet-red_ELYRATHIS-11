package main

import (
    "fmt"
    "github.com/loloboz64/projet-red_ELYRATHIS-11/TF"
)

func main() {
	goblin := TPT.InitGoblin()

	goblin.DisplayHP()

	damage := goblin.CalculateDamage()
	println("Dégâts infligés par le gobelin :", damage)

	goblin.DisplayHP()
}
