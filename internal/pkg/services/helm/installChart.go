package helm

import (
	"errors"
	"fmt"
)

func (s *service) InstallChart(name, releaseName string) error {
	fmt.Println("Entering AddChart with name: ", name)

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
