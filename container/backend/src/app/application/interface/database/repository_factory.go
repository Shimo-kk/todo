package database

import "todo/app/domain/user"

type IRepositoryFactory interface {
	GetUserRepository() user.IUserRepository
}
