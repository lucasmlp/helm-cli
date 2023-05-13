package mongo

import (
	"context"
	"fmt"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (a *adapter) AddRepository(repository *serviceModels.HelmRepository) error {
	fmt.Println("Entering AddRepository with repository: ", repository)

	// Accessing the database
	db := a.client.Database("helm-cli")
	collection := db.Collection("repositories")

	_, err := collection.InsertOne(context.Background(), repository)
	if err != nil {
		return err
	}
	return nil
}
