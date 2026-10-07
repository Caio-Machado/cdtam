package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "monitors",
		Aliases: []string{"list", "ls"},
		Short:   "List available monitors",
		Long:    "List all detected monitors with their ID, resolution and position.",
		Example: "  cdtam list",
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), "list all detected monitors")
			return nil
		},
	}

	return listCmd
}
