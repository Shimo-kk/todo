package usecase

import "todo/app/application/schema"

type IPriorityUsecase interface {
	GetAllPriority() (*[]schema.PriorityReadModel, error)
}
