package main

import (
	"github.com/lucasmlp/helm-cli/internal/pkg/cli"
	"github.com/lucasmlp/helm-cli/internal/pkg/services/helm"
)

const (
	helmRepository = "https://charts.helm.sh/stable"
)

func main() {
	helmService := helm.NewService([]string{})

	helmService.AddRepository(helmRepository)

	cli := cli.NewCLI(helmService)

	cli.Run()
}
