package domain

import "time"

const dateLayout = "2006-01-02"

type Visit struct {
	Fractions []Fraction
	Date      time.Time
}

func NewVisit(date string, fractions []Fraction) (Visit, error) {
	t, err := time.Parse(dateLayout, date)
	if err != nil {
		return Visit{}, err
	}
	return Visit{
		Fractions: fractions,
		Date:      t,
	}, nil
}

func (v Visit) WeightOf(fractionType FractionType) FractionWeight {
	var weight FractionWeight
	for _, f := range v.Fractions {
		if f.Type == fractionType {
			weight += f.Kg
		}
	}
	return weight
}
