package Economie

import "fmt"

var money int = 100
var Market2Unlocked int = 1

func Argent() int {
	return money
}

type Item_market struct {
	Name   string
	Price  int
	Type   string
	Effect string
	Slot   string
}

var Market = []Item_market{
	{Name: "cuirasse", Price: 15, Type: "équipement", Effect: "", Slot: "torse"},
	{Name: "Rubis", Price: 50, Type: "consommable", Effect: "", Slot: ""},
	{Name: "Potion de vie", Price: 10, Type: "consommable", Effect: "", Slot: ""},
	{Name: "Potion de poison", Price: 10, Type: "consommable", Effect: "", Slot: ""},
	{Name: "livre oculte", Price: 100, Type: "consommable", Effect: "", Slot: ""},
	{Name: "Composant Inconus", Price: 999, Type: "???", Effect: "???", Slot: ""},
}

var Market2 = []Item_market{
	{Name: "casque", Price: 15, Type: "équipement", Effect: "", Slot: "tête"},
	{Name: "casque renforcé", Price: 155, Type: "équipement", Effect: "", Slot: "tête"},
	{Name: "casque légendaire", Price: 1555, Type: "équipement", Effect: "", Slot: "tête"},
}

func GetPrice(name string) (int, bool) {
	for _, it := range Market {
		if it.Name == name {
			return it.Price, true
		}
	}
	for i := 0; i < Market2Unlocked && i < len(Market2); i++ {
		if Market2[i].Name == name {
			return Market2[i].Price, true
		}
	}
	return 0, false
}

func Buy(itemName string) string {
	price, found := GetPrice(itemName)
	if !found {
		return "L'objet n'existe pas sur le marché."
	}

	if money >= price {
		money -= price

		if itemName == "Rubis" {
			Market2Unlocked = len(Market2)
		}

		return "Achat réussi de " + itemName + ". Il vous reste " + fmt.Sprint(money) + " pièces."
	}

	return "Fonds insuffisants pour acheter " + itemName + "."
}
