package handler

import (
	cmd "fastshop.com.br/create_pipelines/cmd/interfaces"
	constants "fastshop.com.br/create_pipelines/domain/Constants"
	"github.com/spf13/cobra"
)

// Handler my cobra handler
type Handler struct {
	main     *cobra.Command
	Commands [constants.NumberCommands]cmd.ICmd
}

func (handler *Handler) provideRunInit(cmd cmd.ICmd) {
	cmd.Init()
}

func (handler *Handler) countCommands() int {
	return len(handler.Commands)
}

func (handler *Handler) hasCommand() bool {
	return handler.countCommands() > 0
}

// AddCommandToRoot create new commands
func (handler *Handler) AddCommandToRoot() *Handler {
	if !handler.hasCommand() {
		return handler
	}

	for _, cmd := range handler.Commands {
		handler.provideRunInit(cmd)
		handler.main.AddCommand(cmd.GetCobraCommand())
	}

	return handler

}

// NewCommand Add new command
func (handler *Handler) NewCommand(cmd [constants.NumberCommands]cmd.ICmd) *Handler {
	handler.Commands = cmd
	return handler
}

// Start Command
func (handler *Handler) Start() error {
	return handler.main.Execute()
}

// NewHandler Create my cmd handler
func NewHandler() *Handler {
	return &Handler{
		main: &cobra.Command{
			Use:   "fastshop",
			Short: "do some stuff",
			Long:  "do some stuff",
		},
	}
}
