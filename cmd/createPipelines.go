package cmd

import (
	"errors"

	menus "fastshop.com.br/create_pipelines/cmd/Menus"

	services "fastshop.com.br/create_pipelines/domain/services/pipelines"

	commom "fastshop.com.br/create_pipelines/cmd/commom"

	"github.com/spf13/cobra"
)

// CreatePipelineCmd Command to create pipelines
type CreatePipelineCmd struct {
	common       *commom.CommonCmd
	pipelineMenu *menus.PipelinesMenu
}

// Init init for handler provider
func (c *CreatePipelineCmd) Init() {
	c.common.SetArgs(c.CheckArgs)
	c.common.SetRun(c.Run)
}

// Run my run command
func (c *CreatePipelineCmd) Run(cmd *cobra.Command, args []string) {
	c.pipelineMenu.Create().MountMenu().RunMenu(args)
}

// CheckArgs my check args command
func (c *CreatePipelineCmd) CheckArgs(cmd *cobra.Command, args []string) error {
	if len(args) < 2 {
		return errors.New("numbers of arguments wrong")
	}

	return nil
}

// GetCobraCommand command
func (c *CreatePipelineCmd) GetCobraCommand() *cobra.Command {
	return c.common.GetCobra()
}

// NewCreatePipelineCmd Need to IoC
func NewCreatePipelineCmd(_common *commom.CommonCmd, _createPipelineServices *services.CreatePipelineService, _pipelineMenu *menus.PipelinesMenu) *CreatePipelineCmd {
	return &CreatePipelineCmd{
		common:       _common.CreateCobra("pipelines", "the pipelines commands for fastshop cli", "the pipelines commands for fastshop cli"),
		pipelineMenu: _pipelineMenu,
	}
}
