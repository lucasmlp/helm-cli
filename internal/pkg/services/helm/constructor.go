package helm

import (
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
)

type service struct {
	repositoryList []string
	storageAdapter storage.Adapter
	helmAdapter    helm.Adapter
}

func NewService(
	repositoryList []string,
	storageAdapter storage.Adapter,
	helmAdapter helm.Adapter,
) Service {
	return &service{
		repositoryList: repositoryList,
		storageAdapter: storageAdapter,
		helmAdapter:    helmAdapter,
	}
}
