package helm

import "fmt"

func (s *service) AddChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)

	var found bool
	var err error
	repositoryList := s.storageAdapter.GetRepositoryList()

	for _, repository := range repositoryList {
		fmt.Println("Searching for chart in repo: ", repository)
		found, err = s.helmAdapter.LocateChartInWebRepository(name, repository)
		if err != nil {
			return err
		}

		if found {
			fmt.Println("Found chart in repo: ", repository)
			break
		}
	}

	if !found {
		fmt.Println("Chart not found in any repo")
		return nil
	}

	err = s.storageAdapter.AddChart(name)
	if err != nil {
		return err
	}

	return nil
}
