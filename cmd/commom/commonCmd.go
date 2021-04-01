package commom

import (
	"fastshop.com.br/create_pipelines/infra/helpers"
	"github.com/spf13/cobra"
)

// CommonCmd common struct of cmd
type CommonCmd struct {
	funcHandler *helpers.FuncHandler
	cobra       *cobra.Command
}

// SetRun is a command to set the run command
func (c *CommonCmd) SetRun(f func(cmd *cobra.Command, args []string)) {
	c.cobra.Run = f
}

// SetArgs is a command to set the args command
func (c *CommonCmd) SetArgs(f func(cmd *cobra.Command, args []string) error) {
	c.cobra.Args = f
}

// CreateCobra Create Cobra
func (c *CommonCmd) CreateCobra(command string, short string, long string) *CommonCmd {
	cobra := &cobra.Command{
		Use:   command,
		Short: short,
		Long:  long,
	}

	c.SetCobra(cobra)

	return c
}

// SetCobra provider
func (c *CommonCmd) SetCobra(cobra *cobra.Command) {
	c.cobra = cobra
}

// GetCobra provider
func (c *CommonCmd) GetCobra() *cobra.Command {
	return c.cobra
}

// NewCommonCmd Ioc
func NewCommonCmd(_funcHandler *helpers.FuncHandler) *CommonCmd {
	return &CommonCmd{
		funcHandler: _funcHandler,
	}
}
