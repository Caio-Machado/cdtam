package cli

import (
	"fmt"

	"github.com/caio-machado/cdtam/internal/output"
	"github.com/spf13/cobra"
)

func NewCaptureCmd() *cobra.Command {
	opts := &output.OutputOpts{}

	captureCmd := &cobra.Command{
		Use:   "capture",
		Short: "Capture the screen or a specific monitor",
		Args:  cobra.NoArgs,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if cmd.Flags().Changed("output") && opts.OutDir == "" {
				return fmt.Errorf("--output requires a non-empty value")
			}
			if opts.OutDir == "" && !opts.Clipboard {
				return fmt.Errorf("at least one of the flags --output or --clipboard must be used")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("capture requires a subcommand (available: screen)")
		},
	}

	captureCmd.PersistentFlags().StringVarP(&opts.OutDir, "output", "o", "", "output file or directory (default: ./screenshot-<timestamp>.png)")
	captureCmd.PersistentFlags().BoolVarP(&opts.Clipboard, "clipboard", "c", false, "also copy the capture to the clipboard")

	captureCmd.AddCommand(NewCaptureScreenCmd(opts))

	return captureCmd
}
