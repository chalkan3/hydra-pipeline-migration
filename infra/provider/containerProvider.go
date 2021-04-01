package provider

import (
	"fastshop.com.br/create_pipelines/cmd"
	constants "fastshop.com.br/create_pipelines/domain/Constants"
	"fastshop.com.br/create_pipelines/infra/handler"
	cmdprovider "fastshop.com.br/create_pipelines/infra/provider/cmd"
	entitiesprovider "fastshop.com.br/create_pipelines/infra/provider/domain"
	helpersprovider "fastshop.com.br/create_pipelines/infra/provider/helpers"
	infraprovider "fastshop.com.br/create_pipelines/infra/provider/infra"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	menuprovider "fastshop.com.br/create_pipelines/infra/provider/menuProvider"
	servicesprovider "fastshop.com.br/create_pipelines/infra/provider/services"
	"go.uber.org/dig"
)

// Provider  IoC provider
type Provider struct {
	Container *dig.Container
	providers [constants.NumberProviders]provideinterfaces.IProvider
}

// Provide My Ioc
func (provider *Provider) Provide() *dig.Container {
	for _, pro := range provider.providers {
		pro.Provide(provider.Container)
	}

	return provider.Container
}

// Run Start CLI
func (provider *Provider) Run() error {
	return provider.Provide().Invoke(func(handler *handler.Handler, cmds *cmd.Cmds) {
		handler.NewCommand(cmds.ListCmd()).AddCommandToRoot().Start()
	})
}

func buildContainerProvider() *dig.Container {
	return dig.New()
}

// NewIocProvider create my IoC provider
func NewIocProvider() *Provider {
	return &Provider{
		Container: buildContainerProvider(),
		providers: [constants.NumberProviders]provideinterfaces.IProvider{
			entitiesprovider.NewEntitiesProvider(),
			servicesprovider.NewServiceProvider(),
			cmdprovider.NewCmdProvider(),
			helpersprovider.NewHelpersProvider(),
			menuprovider.NewMenuProvider(),
			infraprovider.NewInfraProvider(),
		},
	}
}
