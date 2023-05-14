package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (a *adapter) GetChartList() ([]*serviceModels.HelmChart, error) {
	db := a.client.Database("helm-cli")
	collection := db.Collection("charts")

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var charts []*serviceModels.HelmChart
	for cursor.Next(context.Background()) {
		var chart *serviceModels.HelmChart
		err := cursor.Decode(&chart)
		if err != nil {
			return nil, err
		}
		charts = append(charts, chart)
	}

	return charts, nil
}
