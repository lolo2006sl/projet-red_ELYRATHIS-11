package TourparTour

type Monster struct {
	Name string
	PV   int
	Atk  int
	Def  int
}

func (m *Monster) CalculateDamage(def int) int {
	damage := m.Atk - def
	if damage < 0 {
		return 0
	}
	return damage
}

func InitGoblin() Monster {
	return Monster{
		Name: "Gobelin",
		PV:   80,
		Atk:  15,
		Def:  3,
	}
}
