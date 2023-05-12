package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func (c cli) addHelmChart(cmd *cobra.Command, args []string) {
	chartName := args[0]
	c.helmService.AddChart(chartName)
}

func (c cli) addIndex(cmd *cobra.Command, args []string) {
	c.helmService.AddIndex()
}

func (c cli) listImages(cmd *cobra.Command, args []string) {
	containerImages, err := c.helmService.ListContainerImages()
	if err != nil {
		log.Fatal(err)
	}

	for _, image := range *containerImages {
		log.Println(image)
	}
}
