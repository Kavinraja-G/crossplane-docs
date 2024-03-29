package cmd

import (
	"github.com/Kavinraja-G/crossplane-docs/pkg"
	"github.com/spf13/cobra"
)

var outputFileName string

// NewCmdMarkdown drives the markdown command for x-docs
func NewCmdMarkdown() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "markdown [PATH]",
		Short:   "Generates markdown based docs for Crossplane resources",
		Aliases: []string{"md"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return pkg.GenMarkdownDocs(cmd, args[0])
		},
	}

	// flags
	cmd.Flags().StringVarP(&outputFileName, "output-file", "o", "xDocs.md", "Filename used for markdown docs output")

	return cmd
}
