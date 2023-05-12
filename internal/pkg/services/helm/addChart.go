package helm

import "fmt"

func (s *service) AddChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)

	var found bool
	var err error
	repositoryList := s.storageAdapter.GetRepositoryList()

	fmt.Printf("repositoryList: %v\n", repositoryList)

	for _, repository := range repositoryList {
		if repository.Local {
			fmt.Println("Searching for chart in repo: ", repository)
			found, err = s.helmAdapter.LocateChartInLocalRepository(name, repository.Location)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("Searching for chart in repo: ", repository)
			found, err = s.helmAdapter.LocateChartInWebRepository(name, repository.Location)
			if err != nil {
				return err
			}
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
