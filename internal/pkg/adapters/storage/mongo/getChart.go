package mongo

import (
	"context"
	"fmt"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *adapter) GetChart(name string) (*models.HelmChart, error) {
	fmt.Println("Entering GetChart with name: ", name)

	db := a.client.Database("helm-cli")
	collection := db.Collection("charts")

	var chart *models.HelmChart
	err := collection.FindOne(context.Background(), bson.M{"name": name}).Decode(&chart)
	if err != nil {
		return nil, err
	}
	return chart, nil
}
