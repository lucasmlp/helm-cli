package helm

func (s service) ListContainerImages() (*[]string, error) {
	return s.helmAdapter.RetrieveContainerImages("./charts")
}
