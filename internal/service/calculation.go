package service

import "github.com/art-es/example-clean-microservice/internal/domain/calculator"

type CalculationService struct {
	Calculator *calculator.Calculator
}

func (svc *CalculationService) Doubling(num float64) float64 {
	// cache ...

	// metric ...

	// another logic ...

	return svc.Calculator.Doubling(num)
}
