package usecase

import (
	"todo/app/application/interface/database"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
)

type priorityUsecase struct {
	databaseHandller database.IDatabaseHandller
}

// 優先度ユースケースの作成
func NewPriorityUsecase(databaseHandller database.IDatabaseHandller) usecase.IPriorityUsecase {
	return &priorityUsecase{databaseHandller: databaseHandller}
}

// 優先度を全件取得
func (u *priorityUsecase) GetAllPriority() (*[]schema.PriorityReadModel, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	priorityRepository := repositoryFactory.GetPriorityRepository()

	// 優先度を全件取得
	priorityEntityList, err := priorityRepository.FindAll()
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := []schema.PriorityReadModel{}
	for _, entity := range *priorityEntityList {
		model := schema.PriorityReadModel{
			Id:        entity.GetId(),
			CreatedAt: entity.GetCreatedAt(),
			UpdatedAt: entity.GetUpdatedAt(),
			Name:      entity.GetName(),
		}
		result = append(result, model)
	}

	return &result, nil
}
