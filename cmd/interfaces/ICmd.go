package cmd

import "github.com/spf13/cobra"

// ICmd Cmd interface
type ICmd interface {
	Init()
	Run(cmd *cobra.Command, args []string)
	CheckArgs(cmd *cobra.Command, args []string) error
	GetCobraCommand() *cobra.Command
}
