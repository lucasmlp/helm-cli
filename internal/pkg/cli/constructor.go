package cli

import "github.com/lucasmlp/helm-cli/internal/pkg/services/helm"

type cli struct {
	helmService helm.Service
}

func NewCLI(helmService helm.Service) cli {
	return cli{
		helmService: helmService,
	}
}
