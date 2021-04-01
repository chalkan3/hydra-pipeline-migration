package infraprovider

import (
	"fastshop.com.br/create_pipelines/infra/handler"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	"go.uber.org/dig"
)

// InfraProvider is my infrastructe provider
type InfraProvider struct {
}

// Provide my provide
func (provider *InfraProvider) Provide(container *dig.Container) {
	container.Provide(handler.NewHandler)
}

// NewInfraProvider Ioc
func NewInfraProvider() provideinterfaces.IProvider {
	return &InfraProvider{}
}
