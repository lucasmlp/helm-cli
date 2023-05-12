package storage

import "fmt"

func (a *adapter) AddChart(name string) error {
	fmt.Println("Entering AddChart with name: ", name)

	a.chartList = append(a.chartList, name)

	fmt.Println("Chart added: ", name)

	fmt.Printf("a.chartList: %v\n", a.chartList)

	return nil
}
