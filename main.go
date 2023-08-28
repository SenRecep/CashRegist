package main

import "fmt"

type Item struct {
	Name     string
	Price    float64
	Discount float64
}

type Describable interface {
	Description() string
}

func (item Item) Description() string {
	result := fmt.Sprintf("%s - %.2f TL", item.Name, item.Price)
	if item.Discount > 0 {
		result += fmt.Sprintf(" (%.2f %% indirimle %.2f TL)", CalculateDiscount(item), CalculatePrice(item))
	}
	return result
}

func CalculatePrice(item Item) float64 {
	return item.Price - item.Discount
}
func CalculateDiscount(item Item) float64 {
	return item.Discount / item.Price * 100
}

func TotalPrice(items []Item) (total float64) {
	for _, item := range items {
		total += CalculatePrice(item)
	}
	return
}

func PrintItem(item Describable) {
	fmt.Printf("%Q\n", item)
}

func PrintItems(items []Item) {
	for _, item := range items {
		PrintItem(item)
	}
	totalPrice := TotalPrice(items)
	fmt.Printf("Toplam Fiyat: %.2f", totalPrice)

}

func (item Item) Format(f fmt.State, verb rune) {
	var value any = item
	if verb == 81 { // check if Q passed
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
	var items []Item
	items = append(items, elma)
	items = append(items, portakal)
	PrintItems(items)
}
