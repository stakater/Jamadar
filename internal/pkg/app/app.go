package app

import "github.com/stakater/Jamadaar/internal/pkg/cmd"

// Run runs the command
func Run() error {
	cmd := cmd.NewJamadaarCommand()
	return cmd.Execute()
}
