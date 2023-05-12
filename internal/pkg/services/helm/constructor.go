package helm

import (
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
)

type service struct {
	storageAdapter storage.Adapter
	helmAdapter    helm.Adapter
}

func NewService(
	storageAdapter storage.Adapter,
	helmAdapter helm.Adapter,
) Service {
	return &service{
		storageAdapter: storageAdapter,
		helmAdapter:    helmAdapter,
	}
}
