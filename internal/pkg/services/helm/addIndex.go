package helm

func (s service) AddIndex() error {
	err := s.helmAdapter.GenerateIndexFile("./charts")
	if err != nil {
		return err
	}

	return nil
}
