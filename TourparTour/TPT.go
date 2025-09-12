package TPT

type Monster struct {
	Name  string
	PVMax int
	PV    int
	def   int
	Atk   int
}

func initGoblin() Monster {
	return Monster{
		Name:  "Goblin",
		PVMax: 10,
		PV:    10,
		def:   1,
		Atk:   3,
	}
}
