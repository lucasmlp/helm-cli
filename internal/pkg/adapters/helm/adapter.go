package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

type adapter struct {
}

type Adapter interface {
	LocateChartInWebRepository(name, url string) (bool, error)
}

func NewAdapter() Adapter {

	return &adapter{}
}

func (a *adapter) LocateChartInWebRepository(name, url string) (bool, error) {
	fmt.Printf("Entering LocateChartInWebRepository with name: %s and repository: %s\n", name, url)

	settings := cli.New()

	repository, err := repo.NewChartRepository(&repo.Entry{
		Name: url,
		URL:  url,
	}, getter.All(settings))
	if err != nil {
		return false, err
	}

	indexFilePath, err := repository.DownloadIndexFile()
	if err != nil {
		return false, err
	}

	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return false, err
	}

	has := indexFile.Has(name, "")
	if has {
		fmt.Printf("Found chart %s in repo %s\n", name, url)
		return true, nil
	}

	return false, nil
}
