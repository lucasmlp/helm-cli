package helm

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

	var seededRand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	releasename := storageChart.Name + fmt.Sprintf("%v", seededRand.Int())

	err = s.helmAdapter.InstallChart(releasename, name)
	if err != nil {
		return err
	}

	return nil
}
