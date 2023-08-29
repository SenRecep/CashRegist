package main

import (
	"fmt"
	"strings"
)

// Item struct represents a product with its name, price, and discount.
type Item struct {
	Name     string
	Price    float64
	Discount float64
}

// Items struct is used to hold a collection of items.
type Items []Item

// Describable interface defines a method for providing a description of an item.
type Describable interface {
	Description() string
}

// Description calculates and returns the description of an item.
func (item Item) Description() string {
	description := fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
	if item.Discount > 0 {
		description += fmt.Sprintf(" (%.2f %% indirimle %.2f TL)", CalculateDiscount(item), CalculatePrice(item))
	}
	return description
}

// Format formats the item based on the provided verb.
func (item Item) Format(f fmt.State, verb rune) {
	var value interface{} = item
	if verb == 'Q' { // 'Q' verb is used for description
		value = item.Description()
	}
	_, err := fmt.Fprint(f, value)
	if err != nil {
		return
	}
}

// Description returns a description of the collection of items.
func (items Items) Description() string {
	var descriptions []string
	for _, item := range items {
		descriptions = append(descriptions, fmt.Sprintf("%Q", item))
	}
	return strings.Join(descriptions, "\n")
}

// CalculatePrice calculates the discounted price of an item.
func CalculatePrice(item Item) float64 {
	return item.Price - item.Discount
}

// CalculateDiscount calculates the discount percentage of an item.
func CalculateDiscount(item Item) float64 {
	return (item.Discount / item.Price) * 100
}

// TotalPrice calculates the total price of a list of items.
func TotalPrice(items []Item) (total float64) {
	for _, item := range items {
		total += CalculatePrice(item)
	}
	return
}

func (items Items) Print() {
	fmt.Println(items.Description())
	totalPrice := TotalPrice(items)
	fmt.Printf("Toplam Fiyat: %.2f", totalPrice)
}

func main() {
	// Sample items
	elma := Item{
		Name:     "Elma",
		Price:    0.75,
		Discount: 0.07,
	}
	portakal := Item{
		Name:     "Portakal",
		Price:    0.75,
		Discount: 0,
	}

	// Create a collection of items
	items := Items{
		elma,
		portakal,
	}
	// Print descriptions of items and calculate total price
	items.Print()
}
