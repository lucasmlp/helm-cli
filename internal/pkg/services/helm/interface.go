package helm

type Service interface {
	AddChart(name string) error
	AddRepository(name, path string) error
	InstallChart(name, releaseName string) error
	AddIndex() error
	ListContainerImages() (*[]string, error)
}
