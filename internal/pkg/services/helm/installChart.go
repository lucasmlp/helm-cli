package helm

import (
	"errors"
	"fmt"
)

func (s *service) InstallChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)

	storageChart, err := s.storageAdapter.GetChart(name)
	if err != nil {
		return err
	}

	if storageChart == nil {
		return errors.New("chart doesn't exist in storage")
	}

	releasename := ""

	err = s.helmAdapter.InstallChart(releasename, name)
	if err != nil {
		return err
	}

	return nil
}
