package domain

import "testing"

func TestCalculatePrice(t *testing.T) {
	testCases := []struct {
		desc      string
		fractions []Fraction
		expected  float64
	}{
		{
			desc: "Green waste 15kg and construction waste 39kg",
			fractions: []Fraction{
				{
					Type: GreenWaste,
					Kg:   15,
				},
				{
					Type: ConstructionWaste,
					Kg:   39,
				},
			},
			expected: 7.35,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			p := NewPrice(tC.fractions)
			result := p.CalculatePrice()
			if result != tC.expected {
				t.Errorf("Expected %f but got %f", tC.expected, result)
			}
		})
	}
}
