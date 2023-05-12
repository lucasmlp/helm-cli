package helm

import (
	"fmt"
	"io"
	"log"
	"os"

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

	fmt.Printf("chartVersion.URLs[0]: %v\n", chartVersion.URLs[0])

	chartCompletePath := path + "/" + name + "-" + chartVersion.Version + ".tgz"

	source, err := os.Open(chartCompletePath)
	if err != nil {
		return nil, err
	}
	defer source.Close()

	destination, err := os.Create("./charts" + "/" + name + "-" + chartVersion.Version + ".tgz")
	if err != nil {
		return nil, err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	if err != nil {
		return nil, err
	}

	err = destination.Sync()
	if err != nil {
		return nil, err
	}

	chartData := &serviceModels.HelmChart{
		Name:        chartVersion.Name,
		Version:     chartVersion.Version,
		Description: chartVersion.Description,
		Path:        path + "/" + name + "-" + chartVersion.Version + ".tgz",
	}

	return chartData, nil
}
