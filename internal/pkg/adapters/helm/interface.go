package helm

import serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"

type Adapter interface {
	LocateChartInWebRepository(name, url string) (*bool, error)
	LocateChartInLocalRepository(name string, path string) (*bool, error)
	GenerateIndexFile(path string) error
	RetrieveContainerImages(repositoryDirectory string) (*[]string, error)
	RetrieveRemoteChart(name, url string) (*serviceModels.HelmChart, error)
	RetrieveLocalChart(name, path string) (*serviceModels.HelmChart, error)
}
