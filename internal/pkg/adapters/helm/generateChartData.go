package helm

import (
	"fmt"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/repo"
)

func generateChartData(path string, name string, chartVersion *repo.ChartVersion) (*serviceModels.HelmChart, error) {

	chart, err := loader.Load(path)
	if err != nil {
		fmt.Println("Error loading chart: ", err)
		return nil, err
	}

	var containerImage string

	switch chart.Values["image"].(type) {
	case string:

		containerImage = chart.Values["image"].(string)

	case map[string]interface{}:

		image := chart.Values["image"].(map[string]interface{})
		imageRepository := image["repository"]
		imageTag := image["tag"]

		if image != nil {
			if imageRepository != nil && imageRepository.(string) != "" {
				containerImage = imageRepository.(string)
				if imageTag != nil && imageTag.(string) != "" {
					containerImage += ":" + imageTag.(string)
				}
			}
		}
	}

	chartData := &serviceModels.HelmChart{
		Name:        chartVersion.Name,
		Version:     chartVersion.Version,
		Description: chartVersion.Description,
		Path:        path,
		Image:       containerImage,
	}

	return chartData, nil
}
