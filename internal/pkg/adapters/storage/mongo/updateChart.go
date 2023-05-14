package mongo

import (
	"context"
	"fmt"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (a *adapter) UpdateChart(chart *serviceModels.HelmChart) error {
	fmt.Println("Entering UpdateChart with name: ", chart.Name)

	db := a.client.Database("helm-cli")
	collection := db.Collection("charts")

	filter := bson.M{"name": chart.Name}
	update := bson.M{"$set": bson.M{"version": chart.Version}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("Error inserting chart: ", err)
		return err
	}

	return nil
}
