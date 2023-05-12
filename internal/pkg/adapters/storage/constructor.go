package storage

type adapter struct {
	repositoryList []string
	chartList      []string
}

func NewAdapter() Adapter {
	return &adapter{}
}
