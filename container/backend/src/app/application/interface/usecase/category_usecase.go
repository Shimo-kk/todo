package usecase

import "todo/app/application/schema"

type ICategoryUsecase interface {
	CreateCategory(userId int, data schema.CategoryCreateModel) error
	UpdateCategory(userId int, data schema.CategoryUpdateModel) error
	GetCategory(userId int, id int) (*schema.CategoryReadModel, error)
	DeleteCategory(userId int, id int) error
	GetAllCategory(userId int) (*[]schema.CategoryReadModel, error)
}
