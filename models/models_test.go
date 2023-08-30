package models_test

import (
	"cashregister/models"
	"fmt"
	"testing"
)

func TestCalculatePrice(t *testing.T) {
	tests := map[string]struct {
		input models.Item
		want  float64
	}{
		"run with normal discount": {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}, want: 80.0},
		"run with full discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 100.0}, want: 0.0},
		"run with over discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 110.0}, want: 0.0},
		"run without discount":     {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 0.0}, want: 100.0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := models.CalculatePrice(&tc.input)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}

func TestCalculateDiscount(t *testing.T) {
	tests := map[string]struct {
		input models.Item
		want  float64
	}{
		"run with normal discount": {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}, want: 20.0},
		"run with zero discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 0.0}, want: 0.0},
		"run with over discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 110.0}, want: 100.0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := models.CalculateDiscount(&tc.input)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}

func TestTotalPrice(t *testing.T) {
	tests := map[string]struct {
		input []models.Item
		want  float64
	}{
		"run zero item":   {input: models.Items{}, want: 0.0},
		"run single item": {input: models.Items{models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}}, want: 80.0},
		"run two item": {input: models.Items{
			models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0},
			models.Item{Name: "Test Item", Price: 100.0, Discount: 0.0}}, want: 180.0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := models.TotalPrice(tc.input)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}

func TestItem_Description(t *testing.T) {
	tests := map[string]struct {
		input models.Item
		want  string
	}{
		"run with normal discount": {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}, want: "Test Item - 100.00 TL (20.00 % indirimle 80.00 TL)"},
		"run with zero discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 0.0}, want: "Test Item - 100.00 TL"},
		"run with full discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 100.0}, want: "Test Item - 100.00 TL (100.00 % indirimle 0.00 TL)"},
		"run with over discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 110.0}, want: "Test Item - 100.00 TL (100.00 % indirimle 0.00 TL)"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.Description()

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}

func TestItems_Description(t *testing.T) {
	tests := map[string]struct {
		input models.Items
		want  string
	}{
		"run zero item":   {input: models.Items{}, want: ""},
		"run single item": {input: models.Items{models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}}, want: "Test Item - 100.00 TL (20.00 % indirimle 80.00 TL)"},
		"run two item": {input: models.Items{
			models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0},
			models.Item{Name: "Test Item", Price: 100.0, Discount: 0.0}}, want: "Test Item - 100.00 TL (20.00 % indirimle 80.00 TL)\nTest Item - 100.00 TL"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.input.Description()

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}

func TestItem_Format(t *testing.T) {
	tests := map[string]struct {
		input models.Item
		want  string
	}{
		"run with normal discount": {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 20.0}, want: "Test Item - 100.00 TL (20.00 % indirimle 80.00 TL)"},
		"run with zero discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 0.0}, want: "Test Item - 100.00 TL"},
		"run with full discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 100.0}, want: "Test Item - 100.00 TL (100.00 % indirimle 0.00 TL)"},
		"run with over discount":   {input: models.Item{Name: "Test Item", Price: 100.0, Discount: 110.0}, want: "Test Item - 100.00 TL (100.00 % indirimle 0.00 TL)"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := fmt.Sprintf("%Q", &tc.input)

			if got != tc.want {
				t.Errorf("want: %v; got: %v", tc.want, got)
			}
		})
	}
}
