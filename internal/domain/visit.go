package domain

import "time"

const dateLayout = "2006-01-02"

type Visit struct {
	Fractions []Fraction
	Date      time.Time
	VisitorID string
	City      City
}

func NewVisit(visitorID string, date string, fractions []Fraction, city City) (Visit, error) {
	t, err := time.Parse(dateLayout, date)
	if err != nil {
		return Visit{}, err
	}
	return Visit{
		Fractions: fractions,
		Date:      t,
		VisitorID: visitorID,
		City:      city,
	}, nil
}
