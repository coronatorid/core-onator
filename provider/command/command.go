package command

import (
	"github.com/spf13/cobra"
)

// Fabricate all command
func Fabricate() *cobra.Command {
	rootCmd := newRoot()
	return rootCmd
}
