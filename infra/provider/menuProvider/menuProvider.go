package menuprovider

import (
	menus "fastshop.com.br/create_pipelines/cmd/Menus"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	"go.uber.org/dig"
)

// MenuProvider is my infrastructe provider
type MenuProvider struct {
}

// Provide my provide
func (provider *MenuProvider) Provide(container *dig.Container) {
	container.Provide(menus.NewPipelinesMenu)
	container.Provide(menus.NewMigrationMenu)
}

// NewMenuProvider Ioc
func NewMenuProvider() provideinterfaces.IProvider {
	return &MenuProvider{}
}
