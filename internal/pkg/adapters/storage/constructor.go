package storage

import serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"

type adapter struct {
	repositoryList []*serviceModels.HelmRepository
	chartList      []*serviceModels.HelmChart
}

func NewAdapter() Adapter {
	return &adapter{}
}
