package interfaces

import "github.com/lakeside763/synapse-crm-pkg/models"

type UserInterface interface {
	CreateUser(data *models.CreateUserInput) (*models.User, error)
}