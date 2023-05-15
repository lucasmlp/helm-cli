package helm

import (
	"errors"
	"fmt"
)

func (s *service) InstallChart(name, releaseName string) error {
	fmt.Println("Installing chart with name ", name, " and release name ", releaseName)

	storageChart, err := s.storageAdapter.GetChart(name)
	if err != nil {
		return errors.New("failed while retrieving chart")
	}

	if storageChart == nil {
		return errors.New("chart doesn't exist in storage")
	}

	err = s.helmAdapter.InstallChart(releaseName, name)
	if err != nil {
		return errors.New("failed while installing chart")
	}

	return nil
}
