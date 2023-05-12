package storage

import serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"

func (a *adapter) GetRepositoryList() []serviceModels.HelmRepository {
	return a.repositoryList
}
