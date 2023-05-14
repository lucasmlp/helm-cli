package helm

import (
	"fmt"

	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) GenerateIndexFile() error {

	index := repo.NewIndexFile()

	chartList, err := a.storageAdapter.GetChartList()
	if err != nil {
		fmt.Println("Error getting chart list from storage adapter")
		return err
	}

	for _, c := range chartList {

		chart, err := loader.Load(c.Path)
		if err != nil {
			fmt.Println("Error loading chart: ", err)
			return err
		}

		err = index.MustAdd(chart.Metadata, c.Name, c.Version, c.Path)
		if err != nil {
			fmt.Println("Error adding chart to index: ", err)
			return err
		}
	}

	index.SortEntries()

	err = index.WriteFile("index.yaml", 0644)
	if err != nil {
		return err
	}

	return nil
}
