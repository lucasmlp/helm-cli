package helm

import (
	"errors"
	"fmt"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (s *service) AddChart(name string) error {
	fmt.Println("Adding chart: ", name)

	var err error
	var chart *models.HelmChart

	storageChart, err := s.storageAdapter.GetChart(name)
	if err != nil {
		return errors.New("failed while retrieving chart")
	}

	if storageChart != nil {
		return errors.New("chart already exists in storage")
	}

	repositoryList, err := s.storageAdapter.GetRepositoryList()
	if err != nil {
		return errors.New("failed while retrieving repository list")
	}

	for _, repository := range repositoryList {
		fmt.Println("Searching for chart in repo: ", repository.Name)

		if repository.Local {
			found, err := s.helmAdapter.LocateChartInLocalRepository(name, repository.Location)
			if err != nil {
				return errors.New("failed while locating chart in local repository")
			}

			if *found {
				chart, err = s.helmAdapter.RetrieveLocalChart(name, repository.Location)
				if err != nil {
					return errors.New("failed while retrieving local chart")
				}
				fmt.Println("Found chart in repo: ", repository.Name)

				err = s.storageAdapter.AddChart(chart)
				if err != nil {
					return errors.New("failed while adding chart to storage")
				}

				return nil
			}
		} else {
			found, err := s.helmAdapter.LocateChartInWebRepository(name, repository.Location)
			if err != nil {
				return errors.New("failed while locating chart in web repository")
			}

			if *found {
				chart, err = s.helmAdapter.RetrieveRemoteChart(name, repository.Location)
				if err != nil {
					fmt.Println("Error retrieving chart: ", err)
					return errors.New("failed while retrieving remote chart")
				}

				fmt.Println("Found chart in repo: ", repository.Name)

				err = s.storageAdapter.AddChart(chart)
				if err != nil {
					return errors.New("failed while adding chart to storage")
				}

				return nil
			}
		}
	}

	return errors.New("chart not found in any repository")
}
