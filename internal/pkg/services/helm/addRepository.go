package helm

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/lucasmlp/helm-cli/internal/pkg/services/models"
)

func (s *service) AddRepository(name, path string) error {
	fmt.Printf("Adding repository with name: %s and path: %s\n", name, path)

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

func (s *service) isValidURL(input string) bool {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return false
	}

	u, err := url.Parse(input)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func (s *service) isValidLocalPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (s *service) addRepository(name, path string, local bool) error {

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
