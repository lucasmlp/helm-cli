package helm

import "fmt"

func (s *service) ListContainerImages() ([]*string, error) {
	fmt.Println("Listing container images")

	chartList, err := s.storageAdapter.GetChartList()
	if err != nil {
		return nil, err
	}

	var containerImages []*string
	for _, chart := range chartList {
		containerImages = append(containerImages, &chart.Image)
	}

	return containerImages, nil
}
