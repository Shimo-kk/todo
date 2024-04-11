package usecase

import (
	"todo/app/application/interface/database"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
	"todo/app/core"
	"todo/app/domain/user"
)

type authUsecase struct {
	databaseHandller database.IDatabaseHandller
}

// 認証ユースケースの作成
func NewAuthUsecase(databaseHandller database.IDatabaseHandller) usecase.IAuthUsecase {
	return &authUsecase{databaseHandller: databaseHandller}
}

// サインアップ
func (u *authUsecase) SignUp(data schema.SignUpModel) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()

		// 同じE-mailアドレスのユーザーが存在しないか確認
		notExists, err := userRepository.NotExists(data.Email)
		if err != nil {
			return err
		}
		if !notExists {
			return core.NewError(core.BadRequestError, "同じE-mailアドレスのユーザーが既に存在しています。")
		}

		// ユーザーのエンティティを作成
		userEntity, err := user.NewEntity(data.Name, data.Email, data.Password)
		if err != nil {
			return err
		}

		// ユーザーを挿入
		_, err = userRepository.Insert(userEntity)
		if err != nil {
			return err
		}

		return nil
	})
}

// サインイン
func (u *authUsecase) SignIn(data schema.SignInModel) (int, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	userRepository := repositoryFactory.GetUserRepository()

	// E-mailアドレスでユーザーを取得
	userEntity, err := userRepository.FindByEmail(data.Email)
	if err != nil {
		return 0, err
	}
	if userEntity == nil {
		return 0, core.NewError(core.NotFoundError, "入力されたE-mailアドレスのユーザーが存在しません。")
	}

	// パスワードを検証
	if err := userEntity.VerifyPassword(data.Password); err != nil {
		return 0, core.NewError(core.BadRequestError, "パスワードに誤りがあります。")
	}

	return userEntity.GetId(), nil
}
