package repository

import (
	"todo/app/core"
	"todo/app/domain/category"

	"gorm.io/gorm"
)

type categoryRepository struct {
	tx *gorm.DB
}

// カテゴリリポジトリの作成
func NewCategoryRepository(tx *gorm.DB) category.ICategoryRepository {
	return &categoryRepository{tx: tx}
}

// 挿入
func (r *categoryRepository) Insert(entity *category.CategoryEntity) (*category.CategoryEntity, error) {
	dto := category.ToDtoFromEntity(entity)
	if err := r.tx.Create(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "カテゴリの作成に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 全件取得
func (r *categoryRepository) FindAll(userId int) (*[]category.CategoryEntity, error) {
	dtos := []category.Category{}
	if err := r.tx.Where(&category.Category{UserId: userId}).Find(&dtos).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "カテゴリの取得に失敗しました。->"+err.Error())
		}
	}
	result := []category.CategoryEntity{}
	for _, dto := range dtos {
		result = append(result, *dto.ToEntity())
	}
	return &result, nil
}

// 主キーで取得
func (r *categoryRepository) FindById(id int) (*category.CategoryEntity, error) {
	dto := category.Category{}
	if err := r.tx.Where(&category.Category{Id: id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "カテゴリの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 更新
func (r *categoryRepository) Update(entity *category.CategoryEntity) (*category.CategoryEntity, error) {
	dto := category.ToDtoFromEntity(entity)
	if err := r.tx.Save(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "カテゴリの更新に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 主キーで削除
func (r *categoryRepository) DeleteById(id int) error {
	if err := r.tx.Delete(&category.Category{Id: id}).Error; err != nil {
		return core.NewError(core.SystemError, "タスクの削除に失敗しました。->"+err.Error())
	}
	return nil
}
