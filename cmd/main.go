package main

import (
	helmAdapter "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
	"github.com/lucasmlp/helm-cli/internal/pkg/cli"
	helmService "github.com/lucasmlp/helm-cli/internal/pkg/services/helm"
)

const (
	helmRepository = "https://charts.helm.sh/stable"
)

func main() {
	storageAdapter := storage.NewAdapter()

	helmAdapter := helmAdapter.NewAdapter()

	repositoryList := []string{helmRepository}

	helmService := helmService.NewService(storageAdapter, helmAdapter)

	for _, repository := range repositoryList {
		helmService.AddRepository(repository)
	}

	cli := cli.NewCLI(helmService)

	cli.Run()
}
