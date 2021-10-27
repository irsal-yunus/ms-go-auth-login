package add

import (
	"github.com/irsal-yunus/ms-go-auth-login/task/data"
	"github.com/irsal-yunus/ms-go-auth-login/task/persistance"
)

type Service struct {
	repository persistance.ITaskDbContext
}

func InitService(repo persistance.ITaskDbContext) *Service {
	return &Service{
		repository: repo,
	}
}

func (service *Service) Add(emp *data.Task) (*data.Task, *data.ErrorDetail) {
	return service.repository.Add(emp)
}
