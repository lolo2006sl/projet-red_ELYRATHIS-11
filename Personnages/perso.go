package hero

// ----- STRUCTURE DU HEROS -----
type Hero struct {
	Name      string
	Classe    string
	PVMax     int
	PV        int
	Def       int
	Atk       int
	Inventory []string
	Wasted    bool // indique si le héros est KO et doit ressusciter au prochain combat
}

// ----- INITIALISATION DES HEROS -----
func InitElise() *Hero {
	return &Hero{
		Name:      "Élise Montclar",
		Classe:    "Érudite",
		PVMax:     45,
		PV:        45,
		Def:       3,
		Atk:       2,
		Inventory: []string{"Livre ancien", "Amulette phocéenne"},
	}
}

func InitJules() *Hero {
	return &Hero{
		Name:      "Jules \"le Noir\" Charvet",
		Classe:    "Mercenaire",
		PVMax:     60,
		PV:        60,
		Def:       5,
		Atk:       4,
		Inventory: []string{"Couteau", "Crochet de fer"},
	}
}

func InitVittorio() *Hero {
	return &Hero{
		Name:      "Dr. Vittorio Santini",
		Classe:    "Médecin-alchimiste",
		PVMax:     50,
		PV:        50,
		Def:       2,
		Atk:       3,
		Inventory: []string{"Flacon de morphine", "Éclat de cristal"},
	}
}
