package cli

import (
	"github.com/spf13/cobra"
)

var (
	verbose bool

	version = "0.1.0"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "cdtam",
		Short: "CLI screenshot tool — capture screens and monitors",
		Long: `cdtam is a simple command-line screenshot tool capable of
	capturing the entire screen or a specific monitor and saving
	the image locally or copying it to the clipboard.`,
		Version:       version,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(NewCaptureCmd())
	rootCmd.AddCommand(NewListCmd())

	return rootCmd
}
