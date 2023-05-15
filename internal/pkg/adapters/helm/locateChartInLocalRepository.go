package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) LocateChartInLocalRepository(name string, path string) (*bool, error) {

	indexFile, err := repo.LoadIndexFile(path + "/index.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	has := indexFile.Has(name, "")

	return &has, nil
}
