package database

import (
	"todo/app/application/interface/database"

	"gorm.io/gorm"
)

type repositoryFactory struct {
	db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) database.IRepositoryFactory {
	return &repositoryFactory{db: db}
}
