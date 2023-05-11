package helm

import (
	"fmt"
	"net/url"
)

func AddRepository(location string) error {
	if isValidURL(location) {
		if err := addHelmRepoFromURL(location); err != nil {
			return err
		}
	} else {
		if err := addHelmRepoFromLocalPath(location); err != nil {
			return err
		}
	}

	fmt.Printf("Helm repo added: %s\n", location)

	return nil
}

func isValidURL(str string) bool {
	_, err := url.ParseRequestURI(str)
	return err == nil
}

func addHelmRepoFromURL(urlStr string) error {
	fmt.Printf("Adding Helm repo from URL: %s\n", urlStr)
	return nil
}

func addHelmRepoFromLocalPath(path string) error {
	fmt.Printf("Adding Helm repo from local path: %s\n", path)
	return nil
}
