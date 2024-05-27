package domain

import (
	"fmt"
	"strings"
)

type FractionPricePerKg float64

const (
	GreenWastePrice        FractionPricePerKg = 0.1
	ConstructionWastePrice FractionPricePerKg = 0.15
)

func (f FractionPricePerKg) Float64() float64 {
	return float64(f)
}

type FractionType string

const (
	GreenWaste        FractionType = "green_waste"
	ConstructionWaste FractionType = "construction_waste"
)

func ParseFractionType(s string) (FractionType, error) {
	trimmedLower := strings.ToLower(strings.TrimSpace(s))
	replaced := strings.Replace(trimmedLower, " ", "_", -1)

	switch replaced {
	case "green_waste":
		return GreenWaste, nil
	case "construction_waste":
		return ConstructionWaste, nil
	default:
		return "", fmt.Errorf("unknown fraction type: %s", s)
	}
}

func (f FractionType) PricePerKg() FractionPricePerKg {
	switch f {
	case GreenWaste:
		return GreenWastePrice
	case ConstructionWaste:
		return ConstructionWastePrice
	default:
		return 0
	}
}

type Fraction struct {
	Type FractionType
	Kg   Weight
}

type Weight float64

func (w Weight) Float64() float64 {
	return float64(w)
}

func ParseWeight(f float64) (Weight, error) {
	// trimmed := strings.TrimSpace(s)
	// f, err := strconv.ParseFloat(trimmed, 64)
	// if err != nil {
	// 	return 0, fmt.Errorf("failed to parse weight: %w", err)
	// }
	if f < 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	return Weight(f), nil
}
