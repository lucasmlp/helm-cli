package mongo

import (
	"context"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *adapter) GetRepository(name string) (*models.HelmRepository, error) {
	db := a.client.Database("helm-cli")
	collection := db.Collection("repositories")

	var repository *models.HelmRepository
	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&repository)
	if err != nil {
		return nil, err
	}
	return repository, nil
}
