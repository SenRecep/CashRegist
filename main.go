package main

import (
	"cashregister/models"
)

func main() {
	// Sample items
	elma := models.Item{
		Name:     "Elma",
		Price:    0.75,
		Discount: 0.07,
	}
	portakal := models.Item{
		Name:     "Portakal",
		Price:    0.75,
		Discount: 0,
	}

	// Create a collection of items
	items := models.Items{
		elma,
		portakal,
	}
	// Print descriptions of items and calculate total price
	items.Print()
}
