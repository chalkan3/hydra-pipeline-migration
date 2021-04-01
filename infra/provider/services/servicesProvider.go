package servicesprovider

import (
	migrationservice "fastshop.com.br/create_pipelines/domain/services/migration"
	services "fastshop.com.br/create_pipelines/domain/services/pipelines"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	"go.uber.org/dig"
)

// ServiceProvider is a provider for services
type ServiceProvider struct {
}

// Provide is a helper Ioc
func (provider *ServiceProvider) Provide(container *dig.Container) {
	container.Provide(services.NewCreatePipelineService)
	container.Provide(migrationservice.NewMigrationService)
}

// NewServiceProvider IoC
func NewServiceProvider() provideinterfaces.IProvider {
	return &ServiceProvider{}
}
