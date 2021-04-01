package cmd

import (
	cmd "fastshop.com.br/create_pipelines/cmd/interfaces"
	constants "fastshop.com.br/create_pipelines/domain/Constants"
)

// Cmds All cmds
type Cmds struct {
	commands [constants.NumberCommands]cmd.ICmd
}

// ListCmd list of all my commands
func (cmds *Cmds) ListCmd() [constants.NumberCommands]cmd.ICmd {
	return cmds.commands
}

// NewCmds need for container IOc
func NewCmds(_create *CreatePipelineCmd, _migration *MigrationsCmd) *Cmds {
	return &Cmds{
		commands: [constants.NumberCommands]cmd.ICmd{
			_create,
			_migration,
		},
	}
}
