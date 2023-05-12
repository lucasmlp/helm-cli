package storage

type Adapter interface {
	AddRepository(location string) error
}
