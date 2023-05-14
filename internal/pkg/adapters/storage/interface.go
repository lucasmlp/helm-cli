package storage

import (
	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

type Adapter interface {
	AddRepository(repository *serviceModels.HelmRepository) error
	AddChart(chart *serviceModels.HelmChart) error
	GetRepositoryList() ([]*serviceModels.HelmRepository, error)
	GetChart(name string) (*serviceModels.HelmChart, error)
	GetChartList() ([]*serviceModels.HelmChart, error)
	GetRepository(name string) (*serviceModels.HelmRepository, error)
}
