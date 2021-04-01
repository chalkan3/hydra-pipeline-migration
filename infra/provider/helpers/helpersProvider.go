package helpersprovider

import (
	"fastshop.com.br/create_pipelines/infra/helpers"
	helpersmenu "fastshop.com.br/create_pipelines/infra/provider/helpers/menu"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	"go.uber.org/dig"
)

// HelpersProvider is a helper Ioc
type HelpersProvider struct {
}

// Provide is a helper Ioc
func (provider *HelpersProvider) Provide(container *dig.Container) {
	container.Provide(helpers.NewFuncHandler)
	container.Provide(helpersmenu.NewMenu)
}

// NewHelpersProvider is a Ioc
func NewHelpersProvider() provideinterfaces.IProvider {
	return &HelpersProvider{}
}
