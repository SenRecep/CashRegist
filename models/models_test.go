package models_test

import (
	"cashregister/models"
	"testing"
)

func TestCalculatePrice(t *testing.T) {
	item := models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}
	expectedPrice := 80.0

	calculatedPrice := models.CalculatePrice(&item)

	if calculatedPrice != expectedPrice {
		t.Errorf("Expected price: %.2f, but got: %.2f", expectedPrice, calculatedPrice)
	}
}

func TestCalculatePriceWithoutDiscount(t *testing.T) {
	item := models.Item{Name: "Test Item", Price: 100.0, Discount: 0}
	expectedPrice := 100.0

	calculatedPrice := models.CalculatePrice(&item)

	if calculatedPrice != expectedPrice {
		t.Errorf("Expected price: %.2f, but got: %.2f", expectedPrice, calculatedPrice)
	}
}

func TestDescription(t *testing.T) {
	item := models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}
	expectedDescription := "Test Item - 100.00 TL (20.00 % indirimle 80.00 TL)"

	calculatedDescription := item.Description()

	if calculatedDescription != expectedDescription {
		t.Errorf("Expected description: %s, but got: %s", expectedDescription, calculatedDescription)
	}
}
