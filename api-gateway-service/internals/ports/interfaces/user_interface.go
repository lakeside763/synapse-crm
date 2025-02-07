package interfaces

import "github.com/lakeside763/synapse-crm-pkg/models"


type UserInterface interface {
	GetUsers() ([]*models.User, error)
	GetUser(id string) (*models.User, error)
	CreateUser(data models.User) (*models.User, error)
}