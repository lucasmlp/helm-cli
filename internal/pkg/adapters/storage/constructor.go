package storage

type adapter struct {
	repositoryList []string
}

func NewAdapter() Adapter {
	return &adapter{}
}
