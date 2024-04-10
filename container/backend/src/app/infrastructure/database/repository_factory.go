package database

import (
	"todo/app/application/interface/database"
	"todo/app/domain/category"
	"todo/app/domain/priority"
	"todo/app/domain/task"
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

func (rf *repositoryFactory) GetTaskRepository() task.ITaskRepository {
	return repository.NewTaskRepository(rf.db)
}

func (rf *repositoryFactory) GetPriorityRepository() priority.IPriorityRepository {
	return repository.NewPriorityRepository(rf.db)
}

func (rf *repositoryFactory) GetCategoryRepository() category.ICategoryRepository {
	return repository.NewCategoryRepository(rf.db)
}
