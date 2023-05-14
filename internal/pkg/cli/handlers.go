package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c cli) addHelmRepository(cmd *cobra.Command, args []string) {
	name := args[0]
	path := args[1]
	err := c.helmService.AddRepository(name, path)
	if err != nil {
		fmt.Println(err)
	}
}

func (c cli) addHelmChart(cmd *cobra.Command, args []string) {
	chartName := args[0]
	err := c.helmService.AddChart(chartName)
	if err != nil {
		fmt.Println(err)
	}
}

func (c cli) installChart(cmd *cobra.Command, args []string) {
	chartName := args[0]
	releaseName := args[1]
	err := c.helmService.InstallChart(chartName, releaseName)
	if err != nil {
		fmt.Println(err)
	}
}

func (c cli) addIndex(cmd *cobra.Command, args []string) {
	c.helmService.AddIndex()
}

func (c cli) listImages(cmd *cobra.Command, args []string) {
	containerImages, err := c.helmService.ListContainerImages()
	if err != nil {
		fmt.Println(err)
	}

	if containerImages != nil {
		for _, image := range containerImages {
			fmt.Println(*image)
		}
	} else {
		fmt.Println("No images found")
	}
}
