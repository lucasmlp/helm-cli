package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func (c cli) Run() {

	var addRepositoryCmd = &cobra.Command{
		Use:   "add-repo [name] [path]",
		Short: "Add a local or remote Helm Repository",
		Args:  cobra.ExactArgs(2),
		Run:   c.addHelmRepository,
	}

	var addChartCmd = &cobra.Command{
		Use:   "add-chart [chart name]",
		Short: "Add a Helm Chart",
		Args:  cobra.ExactArgs(1),
		Run:   c.addHelmChart,
	}

	var installCmd = &cobra.Command{
		Use:   "install chart [chart name] [release name]",
		Short: "Installs a helm chart",
		Args:  cobra.ExactArgs(2),
		Run:   c.installChart,
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

	rootCmd.AddCommand(addRepositoryCmd)
	rootCmd.AddCommand(addChartCmd)
	rootCmd.AddCommand(installCmd)
	rootCmd.AddCommand(indexCmd)
	rootCmd.AddCommand(imagesCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
