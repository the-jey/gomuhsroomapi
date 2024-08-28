package validation

import (
	"github.com/the-jey/gomushroomapi/models"
)

func CreateMushroomValidation(m *models.Mushroom) string {
	// Validate the 'Name' field
	if m.Name == "" {
		s := "Mushroom 'name' can't be empty ❌"
		return s
	}
	if (len(m.Name) < 3) || (len(m.Name) > 224) {
		s := "Mushroom 'name' must be between 3 and 224 characters ❌"
		return s
	}

	// Validate the 'Origin' field
	if m.Origin == "" {
		s := "Mushroom 'origin' can't be empty ❌"
		return s
	}
	if (len(m.Origin) < 3) || (len(m.Origin) > 224) {
		s := "Mushroom 'origin' must be between 3 and 224 characters ❌"
		return s
	}

	// Validate the 'Strenght' field
	if m.Strenght == "" {
		s := "Mushroom 'strenght' can't be empty ❌"
		return s
	}
	if (len(m.Strenght) < 3) || (len(m.Strenght) > 224) {
		s := "Mushroom 'strenght' must be between 3 and 224 characters ❌"
		return s
	}
	if (m.Strenght != "Weak") && (m.Strenght != "Normal") && (m.Strenght != "Strong") && (m.Strenght != "Delusional") {
		s := "Mushroom 'strenght' must be of type : 'Weak' | 'Normal' | 'Strong' | 'Delusional' ❌"
		return s
	}

	// Validate the 'Price' field
	if m.Price <= 0 {
		s := "Mushroom 'price' can't be null or negative ❌"
		return s
	}

	// Validate the 'Quantity' field
	if m.Quantity <= 0 {
		s := "Mushroom 'quantity' can't be null or negative ❌"
		return s
	}

	return ""
}
