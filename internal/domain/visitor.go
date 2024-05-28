package domain

import (
	"fmt"
	"time"
)

var (
	ErrVisitorNotFound = fmt.Errorf("visitor not found")
)

type VisitorType string

const (
	VisitorTypePrivate  VisitorType = "private"
	VisitorTypeBusiness VisitorType = "business"
)

func ParseVisitorType(s string) (VisitorType, error) {
	switch s {
	case "private":
		return VisitorTypePrivate, nil
	case "business":
		return VisitorTypeBusiness, nil
	default:
		return "", fmt.Errorf("unknown visitor type: %s", s)
	}
}

type VisitorRepository interface {
	Get(id string) (Visitor, error)
	Save(v Visitor) error
	Clear() error
}

type Visitor struct {
	ID           string
	City         City
	VisitCounter VisitCounter
	Type         VisitorType
	Visits       []Visit
}

func NewVisitor(id string, city City, visitorType VisitorType) (Visitor, error) {
	return Visitor{
		ID:           id,
		City:         city,
		VisitCounter: VisitCounter{},
		Type:         visitorType,
		Visits:       make([]Visit, 0),
	}, nil
}

func (v *Visitor) RegisterVisit(visit Visit) {
	v.Visits = append(v.Visits, visit)
	v.VisitCounter = v.VisitCounter.AddVisit(visit.Date)
}

func (v *Visitor) LastVisit() Visit {
	return v.Visits[len(v.Visits)-1]
}

func (v *Visitor) CumulativeWeightOverPeriod(period time.Duration, fracType FractionType) FractionWeight {
	lastVisit := v.Visits[len(v.Visits)-1]
	cumWeight := 0.0
	for _, visit := range v.Visits {
		if visit.Date.Add(period).After(lastVisit.Date) {
			cumWeight += visit.WeightOf(fracType).Float64()
		}
	}
	return FractionWeight(cumWeight)
}

type VisitCounter struct {
	Counter        uint16
	LastVisitMonth time.Month
}

func (vc VisitCounter) AddVisit(date time.Time) VisitCounter {
	month := date.Month()
	if month != vc.LastVisitMonth {
		return VisitCounter{
			Counter:        1,
			LastVisitMonth: month,
		}
	}
	return VisitCounter{
		Counter:        vc.Counter + 1,
		LastVisitMonth: month,
	}
}
