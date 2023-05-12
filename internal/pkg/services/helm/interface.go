package helm

type Service interface {
	AddChart(name string) error
	AddRepository(location string) error
}
