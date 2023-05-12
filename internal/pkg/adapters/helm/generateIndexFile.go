package helm

import (
	"helm.sh/helm/v3/pkg/repo"
)

func (a *adapter) GenerateIndexFile(path string) error {
	index, err := repo.IndexDirectory(path, "")
	if err != nil {
		return err
	}

	index.SortEntries()

	err = index.WriteFile(path+"/index.yaml", 0644)
	if err != nil {
		return err
	}
	return nil
}
