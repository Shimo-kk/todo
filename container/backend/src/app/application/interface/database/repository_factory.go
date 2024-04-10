package database

import (
	"todo/app/domain/category"
	"todo/app/domain/priority"
	"todo/app/domain/task"
	"todo/app/domain/user"
)

type IRepositoryFactory interface {
	GetUserRepository() user.IUserRepository
	GetTaskRepository() task.ITaskRepository
	GetPriorityRepository() priority.IPriorityRepository
	GetCategoryRepository() category.ICategoryRepository
}
