package helm

import "fmt"

func (s service) AddIndex() error {
	fmt.Println("Generating index file")

	err := s.helmAdapter.GenerateIndexFile("./charts")
	if err != nil {
		return err
	}

	return nil
}
