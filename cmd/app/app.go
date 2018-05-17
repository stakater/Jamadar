package app

import "github.com/stakater/Jamadar/internal/pkg/cmd"

// Run runs the command
func Run() error {
	cmd := cmd.NewJamadarCommand()
	return cmd.Execute()
}
