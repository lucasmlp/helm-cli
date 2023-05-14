package main

import (
	"context"
	"os"

	helmAdapter "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	mongoAdapter "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mongo"
	"github.com/lucasmlp/helm-cli/internal/pkg/cli"
	helmService "github.com/lucasmlp/helm-cli/internal/pkg/services/helm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	cartRepositoryPath = "./charts"
	fileMode           = 0755
)

func main() {
	err := os.Mkdir(cartRepositoryPath, fileMode)
	if err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	mongoAdapter := mongoAdapter.NewAdapter(client, client.Database("helm-cli"))

	helmAdapter := helmAdapter.NewAdapter(mongoAdapter, "./charts")

	helmService := helmService.NewService(mongoAdapter, helmAdapter)

	cli := cli.NewCLI(helmService)

	cli.Run()
}
