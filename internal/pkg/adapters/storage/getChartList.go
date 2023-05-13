package storage

import serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"

func (a *adapter) GetChartList() ([]*serviceModels.HelmChart, error) {
	return a.chartList, nil
}
