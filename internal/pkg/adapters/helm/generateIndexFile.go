package helm

import (
	"encoding/json"
	"fmt"

	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) GenerateIndexFile(path string) error {

	index := repo.NewIndexFile()

	chartList, err := a.storageAdapter.GetChartList()
	if err != nil {
		fmt.Println("Error getting chart list from storage adapter")
		return err
	}

	prettyChartList, err := prettyStruct(chartList)
	if err != nil {
		fmt.Println("Error pretty printing chart list")
		return err
	}

	fmt.Println("Chart list: ", prettyChartList)

	for _, c := range chartList {
		prettyChart, err := prettyStruct(c)
		if err != nil {
			fmt.Println("Error pretty printing chart")
			return err
		}

		fmt.Println("Chart: ", prettyChart)

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

func prettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
