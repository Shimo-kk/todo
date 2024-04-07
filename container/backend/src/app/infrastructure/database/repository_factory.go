package database

import (
	"todo/app/application/interface/database"
	"todo/app/domain/user"
	"todo/app/infrastructure/database/repository"

	"gorm.io/gorm"
)

type repositoryFactory struct {
	db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) database.IRepositoryFactory {
	return &repositoryFactory{db: db}
}

func (rf *repositoryFactory) GetUserRepository() user.IUserRepository {
	return repository.NewUserRepository(rf.db)
}
