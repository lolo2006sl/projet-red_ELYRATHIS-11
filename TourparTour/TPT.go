package TourparTour

import "fmt"

// ----- STRUCTURE DU HEROS -----
//a changer surement
type Hero struct {
	Name  string
	PVMax int
	PV    int
	Def   int
	Atk   int
}

func InitFakeHero() Hero {
	return Hero{
		Name:  "Héros",
		PVMax: 50,
		PV:    50,
		Def:   3,
		Atk:   6,
	}
}

func (h Hero) CalculateDamage(targetDef int) int {
	if targetDef <= 0 {
		targetDef = 1
	}
	return int(float64(h.Atk) / float64(targetDef))
}

// ----- STRUCTURE DU MONSTRE -----
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
		PVMax: 40,
		PV:    40,
		Def:   2,
		Atk:   5,
	}
}

// Calcul des dégats
func (m Monster) CalculateDamage(targetDef int) int {
	if targetDef <= 0 {
		targetDef = 1
	}
	return int(float64(m.Atk) / float64(targetDef))
}

func (m Monster) DisplayHP() {
	fmt.Printf("%s : %d / %d PV\n", m.Name, m.PV, m.PVMax)
}
