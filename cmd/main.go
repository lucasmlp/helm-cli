package main

import (
	"context"
	"os"

	helmAdapter "github.com/lucasmlp/helm-cli/internal/pkg/adapters/helm"
	"github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage"
	mongoAdapter "github.com/lucasmlp/helm-cli/internal/pkg/adapters/storage/mongo"
	"github.com/lucasmlp/helm-cli/internal/pkg/cli"
	helmService "github.com/lucasmlp/helm-cli/internal/pkg/services/helm"
	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	cartRepositoryPath = "./charts"
	fileMode           = 0755
)

var helmWebRepository = serviceModels.HelmRepository{
	Location: "https://charts.helm.sh/stable",
	Local:    false,
}

var helmLocalRepository = serviceModels.HelmRepository{
	Location: "/Users/lucas/development/helm-charts",
	Local:    true,
}

func main() {
	err := os.Mkdir(cartRepositoryPath, fileMode)
	if err != nil {
		if !os.IsExist(err) {
			panic(err)
		}
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(context.Background())

	mongoAdapter := mongoAdapter.NewAdapter(client, client.Database("helm-cli"))

	storageAdapter := storage.NewAdapter()

	helmAdapter := helmAdapter.NewAdapter(storageAdapter, "./charts")

	// repositoryList := []serviceModels.HelmRepository{helmWebRepository, helmLocalRepository}

	helmService := helmService.NewService(mongoAdapter, helmAdapter)

	// for _, repository := range repositoryList {
	// 	helmService.AddRepository(repository)
	// }

	cli := cli.NewCLI(helmService)

	cli.Run()
}
