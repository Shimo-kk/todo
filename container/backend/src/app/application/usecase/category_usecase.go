package usecase

import (
	"todo/app/application/interface/database"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
	"todo/app/core"
	"todo/app/domain/category"
)

type categoryUsecase struct {
	databaseHandller database.IDatabaseHandller
}

// カテゴリユースケースの作成
func NewCategoryUsecase(databaseHandller database.IDatabaseHandller) usecase.ICategoryUsecase {
	return &categoryUsecase{databaseHandller: databaseHandller}
}

// カテゴリの作成
func (u *categoryUsecase) CreateCategory(userId int, data schema.CategoryCreateModel) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		categoryRepository := rf.GetCategoryRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// カテゴリのエンティティを作成
		categoryEntity, err := category.NewEntity(userEntity.GetId(), data.Name)
		if err != nil {
			return err
		}

		// カテゴリの挿入
		_, err = categoryRepository.Insert(categoryEntity)
		if err != nil {
			return err
		}

		return nil
	})
}

// カテゴリの更新
func (u *categoryUsecase) UpdateCategory(userId int, data schema.CategoryUpdateModel) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		categoryRepository := rf.GetCategoryRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// カテゴリを取得
		categoryEntity, err := categoryRepository.FindById(data.Id)
		if err != nil {
			return err
		}
		if categoryEntity == nil {
			return core.NewError(core.NotFoundError, "カテゴリが存在しません。")
		}

		// 各種プロパティの変更
		if err := categoryEntity.ChangeName(data.Name); err != nil {
			return err
		}

		// カテゴリの更新
		_, err = categoryRepository.Update(categoryEntity)
		if err != nil {
			return err
		}

		return nil
	})
}

// カテゴリの取得
func (u *categoryUsecase) GetCategory(userId int, id int) (*schema.CategoryReadModel, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	userRepository := repositoryFactory.GetTaskRepository()
	categoryRepository := repositoryFactory.GetCategoryRepository()

	// ユーザーを取得
	userEntity, err := userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
	}

	// カテゴリを取得
	categoryEntity, err := categoryRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if categoryEntity == nil {
		return nil, core.NewError(core.NotFoundError, "カテゴリが存在しません。")
	}

	// スキーマへ変換
	result := schema.CategoryReadModel{
		Id:        categoryEntity.GetId(),
		CreatedAt: categoryEntity.GetCreatedAt(),
		UpdatedAt: categoryEntity.GetUpdatedAt(),
		Name:      categoryEntity.GetName(),
	}

	return &result, err
}

// カテゴリの削除
func (u *categoryUsecase) DeleteCategory(userId int, id int) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		categoryRepository := rf.GetCategoryRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// カテゴリの削除
		if err := categoryRepository.DeleteById(id); err != nil {
			return err
		}

		return nil
	})
}

// カテゴリの取得
func (u *categoryUsecase) GetAllCategory(userId int) (*[]schema.CategoryReadModel, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	userRepository := repositoryFactory.GetTaskRepository()
	categoryRepository := repositoryFactory.GetCategoryRepository()

	// ユーザーを取得
	userEntity, err := userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
	}

	// カテゴリを全件取得
	categoryEntityList, err := categoryRepository.FindAll(userEntity.GetId())
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := []schema.CategoryReadModel{}
	for _, entity := range *categoryEntityList {
		model := schema.CategoryReadModel{
			Id:        entity.GetId(),
			CreatedAt: entity.GetCreatedAt(),
			UpdatedAt: entity.GetUpdatedAt(),
			Name:      entity.GetName(),
		}
		result = append(result, model)
	}

	return &result, nil
}
