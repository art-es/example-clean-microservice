package di

import (
	"log"

	"go.uber.org/dig"
)

type Provider struct {
	Constructor interface{}
	Options     []dig.ProvideOption
}

type Providers []*Provider

func (pp Providers) ProvideAll(c *dig.Container) error {
	for _, p := range pp {
		if err := c.Provide(p.Constructor, p.Options...); err != nil {
			log.Printf("[ERROR] diutil: Providers.ProvideAll: c.Provide: %v\n", err)
			return err
		}
	}
	return nil
}
