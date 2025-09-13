package TourparTour

import "fmt"

type Monster struct {
	Name  string
	PVMax int
	PV    int
	Def   int
	Atk   int
}

func InitGoblin() Monster {
	return Monster{
		Name:  "Goblin",
		PVMax: 10,
		PV:    10,
		Def:   2,
		Atk:   3,
	}
}

// Affiche les PV du monstre
func (m Monster) DisplayHP() {
	fmt.Printf("%s : %d / %d PV\n", m.Name, m.PV, m.PVMax)
}

// Calcul des dégâts
func (m Monster) CalculateDamage() int {
	heroDef := 2 // Défense temporaire du héros
	if heroDef == 0 {
		heroDef = 1
	}
	damage := float64(m.Atk) / float64(heroDef)
	return int(damage)
}
