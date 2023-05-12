package helm

import (
	"fmt"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (s *service) AddChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)

	var err error
	var chart *models.HelmChart
	repositoryList := s.storageAdapter.GetRepositoryList()

	fmt.Printf("repositoryList: %v\n", repositoryList)

	for _, repository := range repositoryList {
		fmt.Println("Searching for chart in repo: ", repository)

		if repository.Local {
			chart, err = s.helmAdapter.LocateChartInLocalRepository(name, repository.Location)
			if err != nil {
				return err
			}
		} else {
			chart, err = s.helmAdapter.LocateChartInWebRepository(name, repository.Location)
			if err != nil {
				return err
			}
		}

		if chart != nil {
			fmt.Println("Found chart in repo: ", repository)
			fmt.Printf("chart: %v\n", chart)
			err = s.storageAdapter.AddChart(chart)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return fmt.Errorf("chart %s not found", name)
}
