package helm

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (s *service) AddRepository(name, path string) error {
	fmt.Printf("Entering AddRepository in helm service with name: %s and path: %s\n", name, path)

	repository, err := s.storageAdapter.GetRepository(name)
	if err != nil && err.Error() != "mongo: no documents in result" {
		return err
	}

	if repository != nil {
		return errors.New("repository already exists")
	}

	if s.isValidLocalPath(path) {
		if err := s.addRepository(name, path, true); err != nil {
			return err
		}

		return nil
	} else if s.isValidURL(path) {
		if err := s.addRepository(name, path, false); err != nil {
			return err
		}

		return nil
	}

	return errors.New("invalid repository location")
}

func (s *service) isValidURL(uri string) bool {
	fmt.Println("Validating URL", uri)

	_, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return false
	}

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

func (s *service) addRepository(name, path string, local bool) error {
	fmt.Printf("Adding repository with name: %s, path: %s and local: %v\n", name, path, local)

	repository := models.HelmRepository{
		Name:     name,
		Location: path,
		Local:    local,
	}

	err := s.storageAdapter.AddRepository(&repository)
	if err != nil {
		return err
	}

	return nil
}
