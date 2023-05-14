package mongo

import (
	"context"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *adapter) GetChart(name string) (*models.HelmChart, error) {
	db := a.client.Database("helm-cli")
	collection := db.Collection("charts")

	var chart *models.HelmChart
	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&chart)
	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil
		}
		return nil, err
	}
	return chart, nil
}
