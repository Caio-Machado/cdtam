package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	monitorID   int
	outputPath  string
	toClipboard bool
)

var captureCmd = &cobra.Command{
	Use:     "capture",
	Aliases: []string{"shot", "take"},
	Short:   "Capture the screen or a specific monitor",
	Long: `Capture the entire screen (default) or a specific monitor
and save it as PNG or copy it to the clipboard.`,
	Example: `  cdtam capture                          # capture all screens
  cdtam capture -m 0 -o screen.png         # capture monitor 0 to file
  cdtam capture -m 1 --clipboard           # capture monitor 1 to clipboard
  cdtam capture -o ./shots/ --clipboard    # save to dir AND clipboard`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintln(cmd.OutOrStdout(), "capture the screen or a specific monitor")
		return nil
	},
}

func init() {
	captureCmd.Flags().IntVarP(&monitorID, "monitor", "m", -1, "monitor ID to capture (-1 = all screens)")
	captureCmd.Flags().StringVarP(&outputPath, "output", "o", "", "output file or directory (default: ./screenshot-<timestamp>.png)")
	captureCmd.Flags().BoolVarP(&toClipboard, "clipboard", "c", false, "also copy the capture to the clipboard")

	rootCmd.AddCommand(captureCmd)
}
