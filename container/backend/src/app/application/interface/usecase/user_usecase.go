package usecase

import "todo/app/application/schema"

type IUserUsecase interface {
	GetUser(id int) (*schema.UserReadModel, error)
}
