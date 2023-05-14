package helm

import "fmt"

func (s service) AddIndex() error {
	fmt.Println("Generating index file")

	err := s.helmAdapter.GenerateIndexFile()
	if err != nil {
		return err
	}

	return nil
}
