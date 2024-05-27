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
}

func NewVisitor(id string, city City, visitorType VisitorType) (Visitor, error) {
	return Visitor{
		ID:           id,
		City:         city,
		VisitCounter: VisitCounter{},
		Type:         visitorType,
	}, nil
}

func (v *Visitor) RegisterVisit(date time.Time) {
	v.VisitCounter = v.VisitCounter.AddVisit(date)
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
