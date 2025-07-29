package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmdRoot := &cobra.Command{
		Use:   "notion",
		Short: "Notion API Server",
		Long:  `A RESTful API server built with Go and Gin framework`,
	}

	commands := func() []*cobra.Command {
		return []*cobra.Command{
			server(),
		}
	}

	for _, item := range commands() {
		cmdRoot.AddCommand(item)
	}

	return cmdRoot
}
