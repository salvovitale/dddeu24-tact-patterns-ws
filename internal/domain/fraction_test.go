package domain

import (
	"testing"
	"time"
)

func TestPriceModelCumulatedWeightOverPeriod(t *testing.T) {
	testCases := []struct {
		desc     string
		visitor  Visitor
		expected Price
	}{
		{
			desc: "test case",
			visitor: Visitor{
				ID:           "1",
				City:         CityOakCity,
				VisitCounter: VisitCounter{},
				Visits: []Visit{
					{
						Date: time.Now().AddDate(0, 0, -1),
						Fractions: []Fraction{
							{
								Type: FractionTypeConstructionWaste,
								Kg:   FractionWeight(600),
							},
						},
					},
					{
						Date: time.Now(),
						Fractions: []Fraction{
							{

								Type: FractionTypeConstructionWaste,
								Kg:   FractionWeight(900),
							},
						},
					},
				},
			},
			expected: Price{
				Value:    0.5,
				Currency: EUR,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			model := CumulatedWeightOverPeriod(FractionWeight(1000), Price{Value: 0.21, Currency: EUR}, Price{Value: 0.29, Currency: EUR}, FractionTypeConstructionWaste, time.Duration(24*time.Hour*356))
			price := model(tC.visitor)
			if price != tC.expected {
				t.Errorf("expected: %v, got: %v", tC.expected, price)
			}
		})
	}
}
