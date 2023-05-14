package helm

import (
	"fmt"
)

func (a *adapter) LocateChartInWebRepository(name, url string) (*bool, error) {

	indexFile, err := a.loadRemoteIndex(url)
	if err != nil {
		fmt.Println("Failed to load remote index file")
		fmt.Println(err)
		return nil, err
	}

	has := indexFile.Has(name, "")

	return &has, nil
}
