package cmd

import (
	menus "fastshop.com.br/create_pipelines/cmd/Menus"
	commom "fastshop.com.br/create_pipelines/cmd/commom"
	"github.com/spf13/cobra"
)

// MigrationsCmd is a migration CMD
type MigrationsCmd struct {
	common *commom.CommonCmd
	menu   *menus.MigrationMenu
}

// Init init for handler provider
func (c *MigrationsCmd) Init() {
	c.common.SetArgs(c.CheckArgs)
	c.common.SetRun(c.Run)
}

// Run my run command
func (c *MigrationsCmd) Run(cmd *cobra.Command, args []string) {
	c.menu.Create().MountMenu().RunMenu(args)
}

// CheckArgs my check args command
func (c *MigrationsCmd) CheckArgs(cmd *cobra.Command, args []string) error {
	return nil
}

// GetCobraCommand command
func (c *MigrationsCmd) GetCobraCommand() *cobra.Command {
	return c.common.GetCobra()
}

// NewMigrationsCmd Ioc
func NewMigrationsCmd(_common *commom.CommonCmd, _menu *menus.MigrationMenu) *MigrationsCmd {
	return &MigrationsCmd{
		common: _common.CreateCobra("migrations", "migrations gitlabYAML", "migrations gitlabYAML"),
		menu:   _menu,
	}
}
