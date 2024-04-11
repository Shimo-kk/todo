package usecase

import (
	"todo/app/application/interface/database"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
	"todo/app/core"
)

type userUsecase struct {
	databaseHandller database.IDatabaseHandller
}

// ユーザーユースケースの作成
func NewUserUsecase(databaseHandller database.IDatabaseHandller) usecase.IUserUsecase {
	return &userUsecase{databaseHandller: databaseHandller}
}

// ユーザー取得
func (u *userUsecase) GetUser(id int) (*schema.UserReadModel, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	userRepository := repositoryFactory.GetUserRepository()

	// ユーザーを取得
	userEntity, err := userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
	}

	// スキーマへ変換
	result := schema.UserReadModel{
		Id:    userEntity.GetId(),
		Name:  userEntity.GetName(),
		Email: userEntity.GetEmail(),
	}

	return &result, nil
}
