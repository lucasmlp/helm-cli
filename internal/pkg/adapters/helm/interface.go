package helm

type Adapter interface {
	AddChart(name, repository string) error
	AddRepository(location string) error
}
