package domain

import (
	"fmt"
	"strings"
)

type City string

const (
	Pineville City = "Pineville"
	OakCity   City = "Oak City"
)

func ParseCity(s string) (City, error) {
	trimmedLower := strings.ToLower(strings.TrimSpace(s))
	replaced := strings.Replace(trimmedLower, " ", "_", -1)
	switch replaced {
	case "pineville":
		return Pineville, nil
	case "oak_city":
		return OakCity, nil
	default:
		return "", fmt.Errorf("unknown city: %s", s)
	}
}

type FractionPricePerKg float64

const (
	GreenWastePinevillePrice        FractionPricePerKg = 0.1
	ConstructionPinevilleWastePrice FractionPricePerKg = 0.15
	GreenWasteOakCityPrice          FractionPricePerKg = 0.08
	ConstructionOakCityWastePrice   FractionPricePerKg = 0.19
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

func (f FractionType) PricePerKg(city City) FractionPricePerKg {
	switch f {
	case GreenWaste:
		switch city {
		case Pineville:
			return GreenWastePinevillePrice
		case OakCity:
			return GreenWasteOakCityPrice
		default:
			return 0
		}
	case ConstructionWaste:
		switch city {
		case Pineville:
			return ConstructionPinevilleWastePrice
		case OakCity:
			return ConstructionOakCityWastePrice
		default:
			return 0
		}
	default:
		return 0
	}
}

type Fraction struct {
	Type FractionType
	Kg   FractionWeight
	City City
}

type FractionWeight float64

func (w FractionWeight) Float64() float64 {
	return float64(w)
}

func ParseWeight(f float64) (FractionWeight, error) {
	if f < 0 {
		return 0, fmt.Errorf("weight must be positive")
	}
	return FractionWeight(f), nil
}
