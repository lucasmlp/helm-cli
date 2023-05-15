package helm

import (
	"fmt"
	"log"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) RetrieveLocalChart(name, path string) (*serviceModels.HelmChart, error) {

	indexFile, err := repo.LoadIndexFile(path + "/index.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	chartVersion, err := indexFile.Get(name, "")
	if err != nil {
		log.Println("Error retrieving chart version")
		log.Fatalln(err)
		return nil, err
	}

	chartCompletePath := path + "/" + name + "-" + chartVersion.Version

	return generateChartData(chartCompletePath, name, chartVersion)
}
