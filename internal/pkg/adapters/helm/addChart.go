package helm

import "fmt"

type adapter struct {
	repositoryList []string
}

func (a adapter) AddChart(name, repository string) {
	fmt.Printf("Entering AddChart() with name: %v, repository: %v\n", name, repository)
}
