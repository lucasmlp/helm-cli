package helm

import (
	"fmt"
	"net/url"
	"os"
)

func (s service) AddRepository(location string) error {
	fmt.Printf("Entering AddRepository with location: %s\n", location)

	if s.isValidURL(location) {
		if err := s.addHelmRepo(location); err != nil {
			return err
		}
	} else if s.isValidLocalPath(location) {

		if err := s.addHelmRepo(location); err != nil {
			return err
		}
	}

	fmt.Printf("Helm repo added: %s\n", location)

	return nil
}

func (s service) isValidURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	return err == nil
}

func (s service) isValidLocalPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}

func (s service) addHelmRepo(location string) error {
	fmt.Printf("Adding Helm repo: %s\n", location)
	return nil
}
