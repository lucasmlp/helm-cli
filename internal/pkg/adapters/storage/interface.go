package storage

type Adapter interface {
	AddRepository(location string) error
	AddChart(name string) error
	GetRepositoryList() []string
}
