package repositories

import (
	"github.com/lakeside763/synapse-crm/api-gateway-service/internals/adapters/database"
	"gorm.io/gorm"
)

type DataRepo struct {
	db *gorm.DB
}

func NewDataRepo(databaseUrl string) (*DataRepo, error) {
	db, err := database.NewPostgresConn(databaseUrl)
	if err != nil {
		return nil, err
	}

	return &DataRepo{db}, nil
}