package helm

import (
	"fmt"
	"net/url"
	"os"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (s *service) AddRepository(repository models.HelmRepository) error {
	fmt.Printf("Entering AddRepository in Helm service with location: %s\n", repository.Location)

	if !repository.Local && s.isValidURL(repository.Location) {
		if err := s.addRepository(repository); err != nil {
			return err
		}
	} else if repository.Local && s.isValidLocalPath(repository.Location) {

		if err := s.addRepository(repository); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) isValidURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	return err == nil
}

func (s *service) isValidLocalPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Path does not exist")
			return false
		}
		panic(fmt.Sprintf("Invalid path: %s", path))
	}
	return true
}

func (s *service) addRepository(repository models.HelmRepository) error {
	err := s.storageAdapter.AddRepository(repository)
	if err != nil {
		return err
	}

	return nil
}
