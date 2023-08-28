package main

import "fmt"

// Item struct represents a product with its name, price, and discount.
type Item struct {
	Name     string
	Price    float64
	Discount float64
}

// Describable interface defines a method for providing a description of an item.
type Describable interface {
	Description() string
}

// Description calculates and returns the description of an item.
func (item Item) Description() string {
	result := fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
	if item.Discount > 0 {
		result += fmt.Sprintf(" (%.2f %% indirimle %.2f TL)", CalculateDiscount(item), CalculatePrice(item))
	}
	return result
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

// PrintItem prints the description of a Describable item.
func PrintItem(item Describable) {
	fmt.Printf("%Q\n", item)
}

// PrintItems prints the description of a list of items and the total price.
func PrintItems(items []Item) {
	for _, item := range items {
		PrintItem(item)
	}
	totalPrice := TotalPrice(items)
	fmt.Printf("Toplam Fiyat: %.2f", totalPrice)
}

// Format formats the item based on the provided verb.
func (item Item) Format(f fmt.State, verb rune) {
	var value interface{} = item
	if verb == 81 { // 'Q' verb is used for description
		value = item.Description()
	}
	_, err := fmt.Fprint(f, value)
	if err != nil {
		return
	}
}

var elma = Item{
	Name:     "Elma",
	Price:    0.75,
	Discount: 0.07,
}

var portakal = Item{
	Name:     "Portakal",
	Price:    0.75,
	Discount: 0,
}

func main() {
	items := []Item{
		elma,
		portakal,
	}
	PrintItems(items)
}
