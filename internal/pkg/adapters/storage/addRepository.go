package storage

import (
	"fmt"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (a *adapter) AddRepository(repository *serviceModels.HelmRepository) error {
	fmt.Printf("Entering AddRepository with location: %s\n", repository.Location)

	a.repositoryList = append(a.repositoryList, repository)

	fmt.Printf("Helm repo added: %v\n", repository)

	fmt.Printf("a.repositoryList: %v\n", a.repositoryList)

	return nil
}
