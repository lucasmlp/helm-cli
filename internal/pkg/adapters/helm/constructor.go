package helm

import (
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
)

type adapter struct {
	storageAdapter      storage.Adapter
	chartRepositoryPath string
}

func NewAdapter(
	storageAdapter storage.Adapter,
	chartRepositoryPath string,
) Adapter {
	return &adapter{
		storageAdapter:      storageAdapter,
		chartRepositoryPath: chartRepositoryPath,
	}
}
