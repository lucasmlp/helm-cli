package helm

import "fmt"

func (s *service) AddChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)
	for _, repo := range s.repositoryList {
		//pull repo index and search for chart with name = name
		fmt.Println("Searching for chart in repo: ", repo)
		// add chart info to local chart list and to storage
		fmt.Println("Adding chart to local list and to storage")
	}
	return nil
}
