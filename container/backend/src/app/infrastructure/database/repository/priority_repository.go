package repository

import (
	"todo/app/core"
	"todo/app/domain/priority"

	"gorm.io/gorm"
)

type priorityRepository struct {
	tx *gorm.DB
}

// 優先度リポジトリの作成
func NewPriorityRepository(tx *gorm.DB) priority.IPriorityRepository {
	return &priorityRepository{tx: tx}
}

// 全件取得
func (r *priorityRepository) FindAll() (*[]priority.PriorityEntity, error) {
	dtos := []priority.Priority{}
	if err := r.tx.Find(&dtos).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "優先度の取得に失敗しました。->"+err.Error())
		}
	}
	result := []priority.PriorityEntity{}
	for _, dto := range dtos {
		result = append(result, *dto.ToEntity())
	}
	return &result, nil
}
