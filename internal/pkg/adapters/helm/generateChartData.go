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

	image := chart.Values["image"].(map[string]interface{})
	imageRepository := image["repository"]
	imageTag := image["tag"]

	var containerImage string
	if image != nil && imageRepository != nil {
		if imageTag != nil {
			containerImage = imageRepository.(string) + ":" + imageTag.(string)
		} else {
			containerImage = imageRepository.(string)
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
