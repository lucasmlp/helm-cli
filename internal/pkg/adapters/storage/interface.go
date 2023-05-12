package storage

import serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"

type Adapter interface {
	AddRepository(repository serviceModels.HelmRepository) error
	AddChart(name string) error
	GetRepositoryList() []serviceModels.HelmRepository
}
