package cmdprovider

import (
	"fastshop.com.br/create_pipelines/cmd"
	"fastshop.com.br/create_pipelines/cmd/commom"
	provideinterfaces "fastshop.com.br/create_pipelines/infra/provider/interfaces"
	"go.uber.org/dig"
)

// CmdProvider is a container for cmds
type CmdProvider struct {
}

// Provide the provider
func (provider *CmdProvider) Provide(container *dig.Container) {
	container.Provide(commom.NewCommonCmd)
	container.Provide(cmd.NewCreatePipelineCmd)
	container.Provide(cmd.NewMigrationsCmd)
	container.Provide(cmd.NewCmds)
}

// NewCmdProvider Ioc
func NewCmdProvider() provideinterfaces.IProvider {
	return &CmdProvider{}
}
