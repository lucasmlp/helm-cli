package helm

import "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"

type service struct {
	repositoryList []string
	storageAdapter storage.Adapter
}

func NewService(
	repositoryList []string,
	storageAdapter storage.Adapter,
) Service {
	return &service{
		repositoryList: repositoryList,
		storageAdapter: storageAdapter,
	}
}
