package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) LocateChartInLocalRepository(name string, path string) (*bool, error) {
	fmt.Printf("Entering LocateChartInLocalRepository with name: %s and repository: %s\n", name, path)

	indexFile, err := repo.LoadIndexFile(path + "/index.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	has := indexFile.Has(name, "")

	return &has, nil
}
