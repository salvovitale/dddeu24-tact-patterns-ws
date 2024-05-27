package domain

import "fmt"

type PriceSvc struct {
	userRepo UserRepository
}

type PriceService interface {
	CalculatePrice(fractions []Fraction, id string) (float64, error)
}

func NewPriceSvc(userRepo UserRepository) *PriceSvc {
	return &PriceSvc{
		userRepo: userRepo,
	}
}

func (p *PriceSvc) CalculatePrice(fractions []Fraction, id string) (float64, error) {
	users, err := p.userRepo.GetAll()
	if err != nil {
		return 0, fmt.Errorf("error getting users from repository: %w", err)
	}

	var user *User
	for _, u := range users {
		if u.ID == id {
			user = &u
			break
		}
	}
	if user == nil {
		return 0, fmt.Errorf("user with id %s not found", id)
	}

	var price float64
	for _, f := range fractions {
		price += f.Kg.Float64() * f.Type.PricePerKg(user.City).Float64()
	}
	return price, nil
}
