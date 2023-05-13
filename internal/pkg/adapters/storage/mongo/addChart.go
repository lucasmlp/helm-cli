package mongo

import (
	"context"
	"fmt"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (a *adapter) AddChart(chart *serviceModels.HelmChart) error {
	fmt.Println("Entering AddChart with name: ", chart.Name)
	// Accessing the database
	db := a.client.Database("helm-cli")
	collection := db.Collection("charts")

	_, err := collection.InsertOne(context.Background(), chart)
	if err != nil {
		fmt.Println("Error inserting chart: ", err)
		return err
	}

	return nil
}
