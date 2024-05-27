package domain

import (
	"fmt"
	"time"
)

var (
	ErrVisitorNotFound = fmt.Errorf("visitor not found")
)

type VisitorRepository interface {
	Get(id string) (Visitor, error)
	Save(v Visitor) error
}

func NewVisitor(visit Visit) (Visitor, error) {
	counter, err := NewVisitorCounter(visit.Date)
	if err != nil {
		return Visitor{}, err
	}
	return Visitor{
		ID:           visit.VisitorID,
		City:         visit.City,
		VisitCounter: counter,
	}, nil
}

func (v *Visitor) RegisterVisit(date time.Time) error {
	counter, err := v.VisitCounter.AddVisit(date)
	if err != nil {
		return err
	}
	v.VisitCounter = counter
	return nil
}

type Visitor struct {
	ID           string
	City         City
	VisitCounter VisitCounter
}

type VisitCounter struct {
	Counter        uint16
	LastVisitMonth time.Month
}

func NewVisitorCounter(date time.Time) (VisitCounter, error) {
	month := date.Month()
	return VisitCounter{
		Counter:        1,
		LastVisitMonth: month,
	}, nil
}

func (vc VisitCounter) AddVisit(date time.Time) (VisitCounter, error) {
	month := date.Month()
	if month != vc.LastVisitMonth {
		return VisitCounter{
			Counter:        1,
			LastVisitMonth: month,
		}, nil
	}
	return VisitCounter{
		Counter:        vc.Counter + 1,
		LastVisitMonth: month,
	}, nil

}
