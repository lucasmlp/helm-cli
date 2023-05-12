package main

import (
	"os"

	helmAdapter "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
	"github.com/lucasmlp/helm-cli/internal/pkg/cli"
	helmService "github.com/lucasmlp/helm-cli/internal/pkg/services/helm"
	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

var helmWebRepository = serviceModels.HelmRepository{
	Location: "https://charts.helm.sh/stable",
	Local:    false,
}

var helmLocalRepository = serviceModels.HelmRepository{
	Location: "/Users/lucas/development/helm-charts",
	Local:    true,
}

func main() {
	err := os.Mkdir("charts", 0755)
	if err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	storageAdapter := storage.NewAdapter()

	helmAdapter := helmAdapter.NewAdapter(storageAdapter)

	repositoryList := []serviceModels.HelmRepository{helmWebRepository, helmLocalRepository}

	helmService := helmService.NewService(storageAdapter, helmAdapter)

	for _, repository := range repositoryList {
		helmService.AddRepository(repository)
	}

	cli := cli.NewCLI(helmService)

	cli.Run()
}
