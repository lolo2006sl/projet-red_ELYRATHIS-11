package Economie

var money int = 100

type Item_market struct {
	Name  string
	Price int
}

var market = []Item_market{
	{Name: "cuirasse"}, {Price: 15},
	{Name: "Rubi"}, {Price: 50},
	{Name: "Composant Inconus", Price: 999},
}

func GetPrice(name string) (int, bool) {
	for _, it := range market {
		if it.Name == name {
			return it.Price, true
		}
	}
	return 0, false
}

//Modifier votre marchand lorsque le joueur choisit les items suivants :
//« Potion de vie » : le joueur perd 3 pièces d’or
//« Potion de poison » : le joueur perd 6 pièces d’or
//« Livre de Sort : Boule de feu » : le joueur perd 25 pièces d’or
//Ajouter au menu de vente du marchand les items suivants :
//« Fourrure de Loup » : le joueur perd 4 pièces d’or
//« Peau de Troll » : le joueur perd 7 pièces d’or
//« Cuir de Sanglier » : le joueur perd 3 pièces d’or
//« Plume de Corbeau » : le joueur perd 1 pièce d’or
//Si le joueur choisit un item, il est ajouté à l’inventaire et le coût en pièce d’or est déduit de sa bourse
//d’argent.

//Ajouter au menu principal le choix « Forgeron ». Lorsque le joueur choisit « Forgeron », il doit arriver
//dans un autre menu à choix qui va lui proposer la liste d’équipements à fabriquer suivante :
//Chapeau de l’aventurier
//Tunique de l’aventurier
//Bottes de l’aventurier
//Si le joueur choisit un objet à fabriquer (et qu’il peut le fabriquer), il perd 5 pièces d’or puis
//l’équipement est ajouté à son inventaire.

//recyclage
//augmentation des pris du marché?
