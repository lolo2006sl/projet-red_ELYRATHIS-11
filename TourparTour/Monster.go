package TourparTour

type Monster struct {
	Name  string
	PV    int
	PVMax int
	Atk   int
	Def   int
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
		Name:  "Gobelin",
		PV:    30,
		PVMax: 30,
		Atk:   3,
		Def:   0,
	}
}
