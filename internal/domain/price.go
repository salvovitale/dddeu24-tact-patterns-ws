package domain

type Currency string

const (
	EUR Currency = "EUR"
)

type Price struct {
	Value    float64
	Currency Currency
}
