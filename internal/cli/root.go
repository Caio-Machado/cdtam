package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	verbose bool

	version = "0.1.0"
)

// rootCmd is the base command for cdtam.
var rootCmd = &cobra.Command{
	Use:   "cdtam",
	Short: "CLI screenshot tool — capture screens and monitors",
	Long: `cdtam is a simple command-line screenshot tool capable of
capturing the entire screen or a specific monitor and saving
the image locally or copying it to the clipboard.`,
	Version:       version,
	SilenceUsage:  true,
	SilenceErrors: true,
}

// Execute runs the root command. Called from main.go.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
