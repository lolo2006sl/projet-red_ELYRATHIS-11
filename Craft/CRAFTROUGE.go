package Craft

type Item_craft struct {
	Name  string
	Name2 string
	Name3 string
}

var CraftItems = []Item_craft{
	{Name: "cuirasse", Name2: "cuire" , Name3: "fils"},
	{Name: "Potion de vie", Name2: "fiole en verre" , Name3: "eau de vie"},
	{Name: "Composant Inconus", Name2: "Pièce détaché" , Name3: "Engrenage endomagé"},
}

//Ajouter au menu principal le choix « Forgeron ». Lorsque le joueur choisit « Forgeron », il doit arriver
//dans un autre menu à choix qui va lui proposer la liste d’équipements à fabriquer suivante :
//Chapeau de l’aventurier
//Tunique de l’aventurier
//Bottes de l’aventurier
//Si le joueur choisit un objet à fabriquer (et qu’il peut le fabriquer), il perd 5 pièces d’or puis
//l’équipement est ajouté à son inventaire.

//recyclage
