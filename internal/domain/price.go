package domain

type Currency string

const (
	EUR Currency = "EUR"
)

type TotalPrice struct {
	Fractions []Fraction
	User      User
	Currency  Currency
}

func NewTotalPrice(fractions []Fraction, User User) *TotalPrice {
	return &TotalPrice{
		Fractions: fractions,
		User:      User,
	}
}

func (p *TotalPrice) CalculatePrice() float64 {
	var price float64
	for _, f := range p.Fractions {
		price += f.Kg.Float64() * f.Type.PricePerKg(p.User.City).Float64()
	}
	return price
}
