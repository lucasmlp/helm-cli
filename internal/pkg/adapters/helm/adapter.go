package helm

import (
	"fmt"
	"log"

	serviceModels "github.com/lucasmlp/helm-cli/internal/pkg/services/models"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

type adapter struct {
}

type Adapter interface {
	LocateChartInWebRepository(name, url string) (*serviceModels.HelmChart, error)
	LocateChartInLocalRepository(name string, path string) (*serviceModels.HelmChart, error)
}

func NewAdapter() Adapter {

	return &adapter{}
}

func (a *adapter) LocateChartInWebRepository(name, url string) (*serviceModels.HelmChart, error) {
	fmt.Printf("Entering LocateChartInWebRepository with name: %s and repository: %s\n", name, url)

	settings := cli.New()

	repository, err := repo.NewChartRepository(&repo.Entry{
		Name: url,
		URL:  url,
	}, getter.All(settings))
	if err != nil {
		return nil, err
	}

	indexFilePath, err := repository.DownloadIndexFile()
	if err != nil {
		return nil, err
	}

	indexFile, err := repo.LoadIndexFile(indexFilePath)
	if err != nil {
		return nil, err
	}

	has := indexFile.Has(name, "")
	if has {
		fmt.Printf("Found chart %s in repo %s\n", name, url)

		chartData, err := a.retrieveChart(name, url, indexFile, settings)
		if err != nil {
			log.Fatalln(err)
			return nil, err

		}

		return chartData, nil
	} else {
		fmt.Printf("Chart %s not found in repo %s\n", name, url)
	}

	return nil, nil
}

func (a *adapter) LocateChartInLocalRepository(name string, path string) (*serviceModels.HelmChart, error) {
	fmt.Printf("Entering LocateChartInLocalRepository with name: %s and repository: %s\n", name, path)

	indexFile, err := repo.LoadIndexFile(path + "/index.yaml")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	has := indexFile.Has(name, "")
	if has {
		fmt.Printf("Found chart %s in repo %s\n", name, path)

		chartVersion, err := indexFile.Get(name, "")
		if err != nil {
			log.Println("Error retrieving chart version")
			log.Fatalln(err)
			return nil, err
		}

		fmt.Printf("chartVersion.URLs[0]: %v\n", chartVersion.URLs[0])

		chartData := &serviceModels.HelmChart{
			Name:        chartVersion.Name,
			Version:     chartVersion.Version,
			Description: chartVersion.Description,
			Path:        path + "/" + name + "-" + chartVersion.Version + ".tgz",
		}

		return chartData, nil
	} else {
		fmt.Printf("Chart %s not found in repo %s\n", name, path)
	}

	return nil, nil
}

func (a *adapter) retrieveChart(name, url string, indexFile *repo.IndexFile, settings *cli.EnvSettings) (*serviceModels.HelmChart, error) {
	chartVersion, err := indexFile.Get(name, "")
	if err != nil {
		log.Println("Error retrieving chart version")
		log.Fatalln(err)
		return nil, err
	}

	actionConfiguration := new(action.Configuration)

	err = actionConfiguration.Init(settings.RESTClientGetter(), "", "", log.Printf)
	if err != nil {
		log.Println("Error initializing actionConfiguration")
		log.Fatalln(err)
		return nil, err
	}

	pullClient := action.NewPullWithOpts(action.WithConfig(actionConfiguration))
	pullClient.Settings = settings
	pullClient.DestDir = "./charts"
	pullClient.Untar = true
	pullClient.UntarDir = "./charts"
	pullClient.RepoURL = url
	pullClient.Version = chartVersion.Version

	_, err = pullClient.Run(name)
	if err != nil {
		log.Println("Error pulling chart")
		log.Fatalln(err)
		return nil, err
	}

	chartData := &serviceModels.HelmChart{
		Name:        chartVersion.Name,
		Version:     chartVersion.Version,
		Description: chartVersion.Description,
		Path:        "./charts" + "/" + name + "-" + chartVersion.Version + ".tgz",
	}

	return chartData, nil
}
