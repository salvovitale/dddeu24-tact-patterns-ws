package domain

type Price struct {
	Fractions []Fraction
}

func NewPrice(fractions []Fraction) *Price {
	return &Price{
		Fractions: fractions,
	}
}

func (p *Price) CalculatePrice() float64 {
	var price float64
	for _, f := range p.Fractions {
		price += f.Kg.Float64() * f.Type.PricePerKg().Float64()
	}
	return price
}
