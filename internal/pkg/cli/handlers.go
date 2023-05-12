package cli

import "github.com/spf13/cobra"

func (c cli) addHelmChart(cmd *cobra.Command, args []string) {
	chartName := args[0]
	c.helmService.AddChart(chartName)
}

func (c cli) addIndex(cmd *cobra.Command, args []string) {
	c.helmService.AddIndex()
}
