package cmd

import (
	"github.com/Kavinraja-G/crossplane-docs/pkg"
	"github.com/spf13/cobra"
)

var (
	outputFileName string
	printXRDOnly   bool
)

// NewCmdMarkdown drives the markdown command for x-docs
func NewCmdMarkdown() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "markdown [PATH]",
		Short: "Generates markdown based docs for Crossplane resources",
		Example: `
# Move to the directory which has XRDs & Compositions 
crossplane-docs markdown ./samples -o samples/README.md
`,
		Aliases: []string{"md"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return pkg.GenMarkdownDocs(cmd, args[0])
		},
	}

	// flags
	cmd.Flags().StringVarP(&outputFileName, "output-file", "o", "xDocs.md", "Filename used for markdown docs output")
	cmd.Flags().BoolVarP(&printXRDOnly, "xrd-only", "", false, "Output only XRD specifications")

	return cmd
}
