package cmd

import "github.com/spf13/cobra"

// NewCmdRoot initializes the root command 'crossplane-docs'
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "crossplane-docs",
		Short:   "crossplane-docs - Docs generator for your crossplane resources",
		Version: "0.1.7",
		RunE: func(c *cobra.Command, args []string) error {
			if err := c.Help(); err != nil {
				return err
			}
			return nil
		},
	}

	// child commands
	cmd.AddCommand(NewCmdMarkdown())

	return cmd
}

// Execute drives the root 'crossplane-docs' command
func Execute() error {
	root := NewCmdRoot()
	if err := root.Execute(); err != nil {
		return err
	}

	return nil
}
