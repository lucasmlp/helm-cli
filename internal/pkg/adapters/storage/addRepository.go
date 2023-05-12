package storage

import "fmt"

func (a *adapter) AddRepository(location string) error {
	fmt.Printf("Entering AddRepository with location: %s\n", location)

	a.repositoryList = append(a.repositoryList, location)

	fmt.Printf("Helm repo added: %s\n", location)

	fmt.Printf("a.repositoryList: %v\n", a.repositoryList)

	return nil
}
