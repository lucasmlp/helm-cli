package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func (c cli) Run() {

	var addCmd = &cobra.Command{
		Use:   "add [chart name]",
		Short: "Add a Helm Chart",
		Args:  cobra.ExactArgs(1),
		Run:   c.addHelmChart,
	}

	var indexCmd = &cobra.Command{
		Use:   "index",
		Short: "Add a Helm repository index",
		Run:   c.addIndex,
	}

	var imagesCmd = &cobra.Command{
		Use:   "images",
		Short: "List container images",
		Run:   c.listImages,
	}

	var rootCmd = &cobra.Command{
		Use:   "cli-app",
		Short: "A CLI application to manage Helm charts",
	}

	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(indexCmd)
	rootCmd.AddCommand(imagesCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
