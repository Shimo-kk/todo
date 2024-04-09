package usecase

import "todo/app/application/schema"

type ITaskUsecase interface {
	CreateTask(userId int, data schema.TaskCreateModel) error
	UpdateTask(userId int, data schema.TaskUpdateModel) error
	GetTask(userId int, id int) (*schema.TaskReadModel, error)
	DeleteTask(userId int, id int) error
	DoneTask(userId int, id int) error
	GetAllTask(userId int) (*[]schema.TaskReadModel, error)
}
