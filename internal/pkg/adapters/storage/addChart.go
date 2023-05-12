package storage

import (
	"encoding/json"
	"fmt"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (a *adapter) AddChart(chart *serviceModels.HelmChart) error {
	fmt.Println("Entering AddChart with name: ", chart.Name)

	a.chartList = append(a.chartList, chart)

	fmt.Println("Chart added: ", chart.Name)

	prettyChartList, err := PrettyStruct(a.chartList)
	if err != nil {
		return err
	}

	fmt.Printf("Chart list: %v\n", prettyChartList)

	return nil
}

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
