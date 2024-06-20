package test

import (
	"delivery-app/pkg"
	"testing"
)

func TestGetTotalPrice(t *testing.T) {
	tests := []struct {
		name      string
		pkg       pkg.Package
		basePrice float64
		expected  float64
	}{
		{"pkg weight and distance non zero", pkg.Package{PkgWeight: 2.0, Distance: 5.0}, 100.0, 145.0},
		{"pkg weight zero and distance non zero", pkg.Package{PkgWeight: 0.0, Distance: 0.0}, 100.0, 100.0},
		{"pkg weight non zero and distance zero", pkg.Package{PkgWeight: 1.5, Distance: 0}, 50.0, 65.0},
		{"base price zero", pkg.Package{PkgWeight: 5.0, Distance: 3.0}, 0.0, 65.0},
	}

	for _, test := range tests {
		result := test.pkg.GetTotalPrice(test.basePrice)
		if result != test.expected {
			t.Errorf("%s - For Package{PkgWeight: %f, Distance: %f} with basePrice %f, expected %f but got %f", test.name, test.pkg.PkgWeight, test.pkg.Distance, test.basePrice, test.expected, result)
		}
	}
}
