package entitiesprovider

import (
	"fastshop.com.br/create_pipelines/domain/pipelines"
	"fastshop.com.br/create_pipelines/infra/file"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	yamlcustom "fastshop.com.br/create_pipelines/infra/yaml"
	"go.uber.org/dig"
)

// EntitiesProvider Entities provider
type EntitiesProvider struct {
}

// Provide is a helper Ioc
func (provider *EntitiesProvider) Provide(container *dig.Container) {
	container.Provide(file.NewFile)
	container.Provide(yamlcustom.NewYamlCustom)
	container.Provide(pipelines.NewDotnetPipeline)
}

// NewEntitiesProvider IoC
func NewEntitiesProvider() provideinterfaces.IProvider {
	return &EntitiesProvider{}
}
