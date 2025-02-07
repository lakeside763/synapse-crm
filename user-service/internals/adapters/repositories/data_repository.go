package repositories

import (
	"github.com/lakeside763/synapse-crm/user-service/internals/adapters/database"
	"github.com/lakeside763/synapse-crm/user-service/internals/ports/interfaces"
	"gorm.io/gorm"
)

type DataRepository struct {
	DB *gorm.DB
	User interfaces.UserInterface
}

func NewDataRepository(databaseUrl string) (*DataRepository, error) {
	db, err := database.NewPostgresConn(databaseUrl)
	if err != nil {
		return nil, err
	}

	return &DataRepository{
		DB: db, 
		User: nil,
	}, nil
}