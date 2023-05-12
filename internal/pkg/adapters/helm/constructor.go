package helm

import (
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
)

type adapter struct {
	storageAdapter storage.Adapter
}

func NewAdapter(
	storageAdapter storage.Adapter,
) Adapter {
	return &adapter{
		storageAdapter: storageAdapter,
	}
}
