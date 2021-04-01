package helpers

import "github.com/spf13/cobra"

// FuncHandler improve the handle of cobra func
type FuncHandler struct {
}

// Run  the run function of cobra
func (h *FuncHandler) Run(run func(cmd *cobra.Command, args []string)) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		run(cmd, args)
	}
}

// CheckArgs the checkargs function of cobra
func (h *FuncHandler) CheckArgs(check func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		return check(cmd, args)
	}
}

// NewFuncHandler Ioc
func NewFuncHandler() *FuncHandler {
	return &FuncHandler{}
}
