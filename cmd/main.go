package main

import (
	"fmt"
	"log"

	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	"github.com/spf13/cobra"
)

const (
	helmRepository = "https://charts.helm.sh/stable"
)

var rootCmd = &cobra.Command{
	Use:   "cli-app",
	Short: "A CLI application to manage Helm charts",
}

var addCmd = &cobra.Command{
	Use:   "add [chart name]",
	Short: "Add a Helm Chart",
	Args:  cobra.ExactArgs(1),
	Run:   addHelmChart,
}

func main() {
	helm.AddRepository(helmRepository)

	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func addHelmChart(cmd *cobra.Command, args []string) {
	chartName := args[0]
	fmt.Printf("chartName: %v\n", chartName)
}
