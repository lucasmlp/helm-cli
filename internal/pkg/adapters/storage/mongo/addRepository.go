package mongo

import (
	"context"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (a *adapter) AddRepository(repository *serviceModels.HelmRepository) error {
	db := a.client.Database("helm-cli")
	collection := db.Collection("repositories")

	_, err := collection.InsertOne(context.Background(), repository)
	if err != nil {
		return err
	}
	return nil
}
