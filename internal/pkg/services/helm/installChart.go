package helm

import (
	"errors"
	"fmt"
)

func (s *service) InstallChart(name, releaseName string) error {
	fmt.Println("Installing chart with name ", name, " and release name ", releaseName)

	storageChart, err := s.storageAdapter.GetChart(name)
	if err != nil {
		return err
	}

	if storageChart == nil {
		return errors.New("chart doesn't exist in storage")
	}

	err = s.helmAdapter.InstallChart(releaseName, name)
	if err != nil {
		return err
	}

	return nil
}
