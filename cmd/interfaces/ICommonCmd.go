package cmd

import (
	"github.com/spf13/cobra"
)

// ICommonCmd common interface of cmds
type ICommonCmd interface {
	SetRun(f func(cmd *cobra.Command, args []string))
	SetArgs(f func(cmd *cobra.Command, args []string) error)
	GetCobra()
	SetCobra(cobra *cobra.Command)
	CreateCobra(command string, short string, long string)
}
