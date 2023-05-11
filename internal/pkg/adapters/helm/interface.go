package helm

type Adapter interface {
	AddRepository(location string) error
}
