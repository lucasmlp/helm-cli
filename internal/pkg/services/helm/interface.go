package helm

import "github.com/lucasmlp/helm-cli/internal/pkg/services/models"

type Service interface {
	AddChart(name string) error
	AddRepository(repository models.HelmRepository) error
}
