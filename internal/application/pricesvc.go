package application

import (
	"fmt"

	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
)

type PriceUseCase interface {
	CalculatePrice(fractions []domain.Fraction, visitorID string, date string) (*domain.Price, error)
	ClearScenario() error
}

type ExtUser struct {
	ID          string
	City        domain.City
	VisitorType domain.VisitorType
}

type extUserService interface {
	Get(id string) (*ExtUser, error)
}

type PriceAppService struct {
	priceSvc       domain.PriceService
	visitorRepo    domain.VisitorRepository
	extUserService extUserService
}

func NewPriceAppService(priceSvc domain.PriceService, visitorRepo domain.VisitorRepository, extUserService extUserService) *PriceAppService {
	return &PriceAppService{
		priceSvc:       priceSvc,
		visitorRepo:    visitorRepo,
		extUserService: extUserService,
	}
}

func (p *PriceAppService) CalculatePrice(fractions []domain.Fraction, visitorID string, date string) (*domain.Price, error) {

	user, err := p.extUserService.Get(visitorID)
	if err != nil {
		return nil, fmt.Errorf("error getting user from external service: %w", err)
	}

	visitor, err := p.visitorRepo.Get(user.ID)
	if err != nil {
		if err == domain.ErrVisitorNotFound {
			visitor, err = domain.NewVisitor(user.ID, user.City, user.VisitorType)
			if err != nil {
				return nil, fmt.Errorf("error creating visitor: %w", err)
			}
		} else {
			return nil, fmt.Errorf("error getting visitor: %w", err)
		}
	}

	visit, err := domain.NewVisit(date, fractions)
	if err != nil {
		return nil, fmt.Errorf("error creating visit: %w", err)
	}

	visitor.RegisterVisit(visit)

	price, err := p.priceSvc.CalculatePrice(visitor)
	if err != nil {
		return nil, fmt.Errorf("error calculating price: %w", err)
	}

	err = p.visitorRepo.Save(visitor)
	if err != nil {
		return nil, fmt.Errorf("error saving visitor: %w", err)
	}

	return &price, nil
}

func (p *PriceAppService) ClearScenario() error {
	return p.visitorRepo.Clear()
}
