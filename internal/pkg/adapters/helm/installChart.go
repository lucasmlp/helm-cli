package helm

import (
	"log"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func (a adapter) InstallChart(releaseName, name string) error {

	storageChart, err := a.storageAdapter.GetChart(name)
	if err != nil {
		log.Println(err)
		return err
	}

	validatedChart, err := loadAndValidate(storageChart.Path)
	if err != nil {
		log.Println(err)
	}

	actionConfiguration := new(action.Configuration)

	settings := cli.New()

	err = actionConfiguration.Init(settings.RESTClientGetter(), "default", "", log.Printf)
	if err != nil {
		log.Println("Error initializing actionConfiguration")
		log.Fatalln(err)
		return err
	}

	client := action.NewInstall(actionConfiguration)
	client.ReleaseName = releaseName
	client.Namespace = "default"

	_, err = client.Run(validatedChart, nil)
	if err != nil {
		log.Println("Error installing chart")
		log.Println(err)
		return err
	}

	return nil
}

func loadAndValidate(chartPath string) (*chart.Chart, error) {

	chart, err := loader.Load(chartPath)
	if err != nil {
		log.Println("Error loading chart: ", err)
		return nil, err
	}

	err = chart.Validate()
	if err != nil {
		log.Println("Error validating chart: ", err)
		return nil, err
	}

	log.Printf("chart %s is valid \n", chart.Metadata.Name)

	return chart, nil
}
