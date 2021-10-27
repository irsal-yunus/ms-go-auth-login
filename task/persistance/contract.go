package persistance

import (
	"github.com/irsal-yunus/ms-go-auth-login/task/data"
)

type ITaskDbContext interface {
	GetAll() ([]data.Task, *data.ErrorDetail)
	GetById(id int) (*data.Task, *data.ErrorDetail)
	Update(emp *data.Task) (*data.Task, *data.ErrorDetail)
	Add(emp *data.Task) (*data.Task, *data.ErrorDetail)
	Delete(id int) (*data.Task, *data.ErrorDetail)
}
