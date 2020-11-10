package command

import (
	"github.com/spf13/cobra"
)

// Command manage all command in coronator
type Command struct {
	rootCmd *cobra.Command
}

// Fabricate root command
func Fabricate() *Command {
	return &Command{
		rootCmd: newRoot(),
	}
}

// SetArgs set argument for command line interface
func (c *Command) SetArgs(args []string) {
	c.rootCmd.SetArgs(args)
}

// Execute command line interface
func (c *Command) Execute() error {
	return c.rootCmd.Execute()
}

// InjectCommand inject new command into command list
func (c *Command) InjectCommand(s Scaffold) {
	c.rootCmd.AddCommand(
		&cobra.Command{
			Use:     s.Use(),
			Short:   s.Short(),
			Example: s.Example(),
			Run: func(cmd *cobra.Command, args []string) {
				s.Run(args)
			},
		},
	)

}
