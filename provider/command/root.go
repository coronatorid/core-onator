package command

import "github.com/spf13/cobra"

type root struct{}

func newRoot() *cobra.Command {
	return &cobra.Command{
		Use:     "core-onator",
		Short:   "Core business logic of coronator",
		Example: "core-onator",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
}
