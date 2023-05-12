package main

import (
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
	"github.com/lucasmlp/helm-cli/internal/pkg/cli"
	"github.com/lucasmlp/helm-cli/internal/pkg/services/helm"
)

const (
	helmRepository = "https://charts.helm.sh/stable"
)

func main() {
	storageAdapter := storage.NewAdapter()
	repositoryList := []string{helmRepository}

	helmService := helm.NewService(repositoryList, storageAdapter)

	cli := cli.NewCLI(helmService)

	cli.Run()
}
