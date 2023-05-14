package helm

type Service interface {
	AddChart(name string) error
	AddRepository(path string) error
	AddIndex() error
	ListContainerImages() (*[]string, error)
}
