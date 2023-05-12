package helm

import "fmt"

func (s *service) AddChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)

	var found bool
	var err error
	for _, repository := range s.repositoryList {
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

	// add chart info to local chart list and to storage
	fmt.Println("Adding chart to local list and to storage")

	return nil
}
