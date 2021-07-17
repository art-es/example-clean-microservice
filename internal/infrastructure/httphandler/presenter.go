package httphandler

import (
	"log"

	"github.com/art-es/example-clean-microservice/internal/domain/calculator"
	"github.com/art-es/example-clean-microservice/internal/presentation"
	"github.com/art-es/example-clean-microservice/internal/service"
	. "github.com/art-es/example-clean-microservice/internal/util/di"

	"go.uber.org/dig"
)

func newPresenter() (pres *presentation.HTTPPresenter, err error) {
	c := dig.New()

	pp := Providers{
		calculatorProvider(),
		calculationServiceProvider(),
	}
	if err = pp.ProvideAll(c); err != nil {
		return
	}

	pres, err = invokeHTTPPresenter(c)
	if err != nil {
		log.Printf("[ERROR] httphandler: newPresenter: c.Invoke(PresenterConstructor): %v\n", err)
	}
	return pres, nil
}

func calculatorProvider() *Provider {
	return &Provider{
		Constructor: func() *calculator.Calculator { return new(calculator.Calculator) },
	}
}

func calculationServiceProvider() *Provider {
	return &Provider{
		Constructor: func(c *calculator.Calculator) *service.CalculationService {
			return &service.CalculationService{Calculator: c}
		},
	}
}

func invokeHTTPPresenter(c *dig.Container) (pres *presentation.HTTPPresenter, err error) {
	err = c.Invoke(func(calculationService *service.CalculationService) {
		pres = &presentation.HTTPPresenter{
			CalculationService: calculationService,
		}
	})
	return
}
