package interfaces

import (
	"CRUD+Check/pkg/models"
)

type IDataWorker interface{
	CreateUser(models.User) error
	GetUser(string) (models.User, error)
	UpdateUser(models.User) error
	DeleteUser(string) error
	CheckUser(models.User) error
}