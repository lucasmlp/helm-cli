package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) loadRemoteIndex(url string) (*repo.IndexFile, error) {
	settings := cli.New()

	repository, err := repo.NewChartRepository(&repo.Entry{
		Name: url,
		URL:  url,
	}, getter.All(settings))

	if err != nil {
		fmt.Println("Failed to create chart repository")
		fmt.Println(err)
		return nil, err
	}

	repository.Config.InsecureSkipTLSverify = true

	indexFilePath, err := repository.DownloadIndexFile()
	if err != nil {
		fmt.Println("Failed to download index file")
		fmt.Println(err)
		return nil, err
	}

	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		fmt.Println("Failed to load index file")
		fmt.Println(err)
		return nil, err
	}

	return indexFile, nil
}
