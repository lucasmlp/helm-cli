package mongo

import (
	"context"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *adapter) GetRepositoryList() ([]*serviceModels.HelmRepository, error) {
	db := a.client.Database("helm-cli")
	collection := db.Collection("repositories")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var repositoryList []*serviceModels.HelmRepository
	for cursor.Next(context.Background()) {
		var repository *serviceModels.HelmRepository
		err := cursor.Decode(&repository)
		if err != nil {
			return nil, err
		}
		repositoryList = append(repositoryList, repository)
	}

	return repositoryList, nil
}
