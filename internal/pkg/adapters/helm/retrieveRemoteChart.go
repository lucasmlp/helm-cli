package helm

import (
	"log"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
)

func (a *adapter) RetrieveRemoteChart(name, url string) (*serviceModels.HelmChart, error) {
	indexFile, err := a.loadRemoteIndex(url)
	if err != nil {
		return nil, err
	}

	chartVersion, err := indexFile.Get(name, "")
	if err != nil {
		log.Println("Error retrieving chart version")
		log.Fatalln(err)
		return nil, err
	}

	actionConfiguration := new(action.Configuration)

	settings := cli.New()

	err = actionConfiguration.Init(settings.RESTClientGetter(), "", "", log.Printf)
	if err != nil {
		log.Println("Error initializing actionConfiguration")
		log.Fatalln(err)
		return nil, err
	}

	pullClient := action.NewPullWithOpts(action.WithConfig(actionConfiguration))
	pullClient.Settings = settings
	pullClient.DestDir = a.chartRepositoryPath
	pullClient.RepoURL = url
	pullClient.Untar = true
	pullClient.Version = chartVersion.Version
	pullClient.InsecureSkipTLSverify = true

	_, err = pullClient.Run(name)
	if err != nil {
		log.Println("Error pulling chart")
		log.Fatalln(err)
		return nil, err
	}

	chartCompletePath := a.chartRepositoryPath + "/" + name

	return generateChartData(chartCompletePath, name, chartVersion)
}
