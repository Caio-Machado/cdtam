package cli

import (
	"github.com/caio-machado/cdtam/internal/capture"
	"github.com/caio-machado/cdtam/internal/output"
	"github.com/spf13/cobra"
)

func NewCaptureScreenCmd(outOpts *output.OutputOpts) *cobra.Command {
	captureScreenCmd := &cobra.Command{
		Use:   "screen",
		Short: "capture the primary screen",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			outputs, err := outOpts.Resolve()
			if err != nil {
				return err
			}

			return capture.Run(capture.NewVirtualScreen(), outputs)
		},
	}

	return captureScreenCmd
}
