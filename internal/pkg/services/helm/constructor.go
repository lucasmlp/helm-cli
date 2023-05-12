package helm

type service struct {
	repositoryList []string
}

func NewService(repositoryList []string) Service {
	return service{
		repositoryList: repositoryList,
	}
}
