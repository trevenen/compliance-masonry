package validate

import (
	"io"

	"github.com/opencontrol/compliance-masonry/validate"
	"github.com/spf13/cobra"
)

// NewCmdValidate validates the current masonry
func NewCmdValidate(out io.Writer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validates the current opencontrol masonry repository. Use get command to create opencontrol masonry repository.",
		Run: func(cmd *cobra.Command, args []string) {
			validate.Validate()
		},
	}
	return cmd
}
