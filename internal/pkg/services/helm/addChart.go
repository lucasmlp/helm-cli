package helm

import (
	"errors"
	"fmt"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (s *service) AddChart(name string) error {
	fmt.Println("Adding chart: ", name)

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
			chartAdded, err := s.findLocalChartAndAdd(name, repository)
			if err != nil {
				return err
			}

			if *chartAdded {
				return nil
			}

		} else {
			chartAdded, err := s.findRemoteChartAndAdd(name, repository)
			if err != nil {
				return err
			}

			if *chartAdded {
				return nil
			}
		}
	}

	return errors.New("chart not found in any repository")
}

func (s *service) findLocalChartAndAdd(name string, localRepository *models.HelmRepository) (*bool, error) {

	found, err := s.helmAdapter.LocateChartInLocalRepository(name, localRepository.Location)
	if err != nil {
		return nil, errors.New("failed while locating chart in local repository")
	}

	chartAdded := false
	if *found {
		chart, err := s.helmAdapter.RetrieveLocalChart(name, localRepository.Location)
		if err != nil {
			return nil, errors.New("failed while retrieving local chart")
		}
		fmt.Println("Found chart in repo: ", localRepository.Name)

		err = s.storageAdapter.AddChart(chart)
		if err != nil {
			return nil, errors.New("failed while adding chart to storage")
		}

		chartAdded = true
	}

	return &chartAdded, nil
}

func (s *service) findRemoteChartAndAdd(name string, remoteRepository *models.HelmRepository) (*bool, error) {
	found, err := s.helmAdapter.LocateChartInWebRepository(name, remoteRepository.Location)
	if err != nil {
		return nil, errors.New("failed while locating chart in remote repository")
	}

	chartAdded := false
	if *found {
		chart, err := s.helmAdapter.RetrieveRemoteChart(name, remoteRepository.Location)
		if err != nil {
			fmt.Println("Error retrieving chart: ", err)
			return nil, errors.New("failed while retrieving remote chart")
		}

		fmt.Println("Found chart in repo: ", remoteRepository.Name)

		err = s.storageAdapter.AddChart(chart)
		if err != nil {
			return nil, errors.New("failed while adding chart to storage")
		}

		chartAdded = true
	}
	return &chartAdded, nil
}
