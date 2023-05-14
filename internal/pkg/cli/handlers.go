package cli

import (
	"log"

	"github.com/spf13/cobra"
)

func (c cli) addHelmRepository(cmd *cobra.Command, args []string) {
	name := args[1]
	path := args[2]
	err := c.helmService.AddRepository(name, path)
	if err != nil {
		log.Fatal(err)
	}
}

func (c cli) addHelmChart(cmd *cobra.Command, args []string) {
	chartName := args[0]
	err := c.helmService.AddChart(chartName)
	if err != nil {
		log.Fatal(err)
	}
}

func (c cli) installChart(cmd *cobra.Command, args []string) {
	chartName := args[1]
	releaseName := args[2]
	err := c.helmService.InstallChart(chartName, releaseName)
	if err != nil {
		log.Fatal(err)
	}
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
