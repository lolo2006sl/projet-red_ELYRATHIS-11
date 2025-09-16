package Craft

type Item_craft struct {
	Name  string
	Name2 string
}

var CraftItems = []Item_craft{
	{Name: "cuirasse", Name2: "a"},
	{Name: "Rubis", Name2: "a"},
	{Name: "Potion de vie", Name2: "a"},
	{Name: "Potion de poison", Name2: "a"},
	{Name: "livre oculte", Name2: ""},
	{Name: "Composant Inconus", Name2: "a"},
}

//Ajouter au menu principal le choix « Forgeron ». Lorsque le joueur choisit « Forgeron », il doit arriver
//dans un autre menu à choix qui va lui proposer la liste d’équipements à fabriquer suivante :
//Chapeau de l’aventurier
//Tunique de l’aventurier
//Bottes de l’aventurier
//Si le joueur choisit un objet à fabriquer (et qu’il peut le fabriquer), il perd 5 pièces d’or puis
//l’équipement est ajouté à son inventaire.

//recyclage
