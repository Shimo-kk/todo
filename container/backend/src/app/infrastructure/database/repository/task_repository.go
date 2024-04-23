package repository

import (
	"todo/app/core"
	"todo/app/domain/task"

	"gorm.io/gorm"
)

type taskRepository struct {
	tx *gorm.DB
}

// タスクリポジトリの作成
func NewTaskRepository(tx *gorm.DB) task.ITaskRepository {
	return &taskRepository{tx: tx}
}

// 挿入
func (r *taskRepository) Insert(entity *task.TaskEntity) (*task.TaskEntity, error) {
	dto := task.ToDtoFromEntity(entity)
	if err := r.tx.Create(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "タスクの作成に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 全件取得
func (r *taskRepository) FindAll(userId int) (*[]task.TaskEntity, error) {
	dtos := []task.Task{}
	if err := r.tx.Where(&task.Task{UserId: userId}).Find(&dtos).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "タスクの取得に失敗しました。->"+err.Error())
		}
	}
	result := []task.TaskEntity{}
	for _, dto := range dtos {
		result = append(result, *dto.ToEntity())
	}
	return &result, nil
}

// IDで取得
func (r *taskRepository) FindById(id int) (*task.TaskEntity, error) {
	dto := task.Task{}
	if err := r.tx.Where(&task.Task{Id: id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "タスクの取得に失敗しました。->"+err.Error())
		}
	}
	result := dto.ToEntity()
	return result, nil
}

// 更新
func (r *taskRepository) Update(entity *task.TaskEntity) (*task.TaskEntity, error) {
	dto := task.ToDtoFromEntity(entity)
	old := task.Task{}
	if err := r.tx.Where(&task.Task{Id: dto.Id}).First(&dto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		} else {
			return nil, core.NewError(core.SystemError, "タスクの取得に失敗しました。->"+err.Error())
		}
	}

	if !dto.UpdatedAt.After(old.UpdatedAt) {
		return nil, core.NewError(core.ConflictError, "タスクの更新に失敗しました。別のリクエストが同じリソースを変更しています。")
	}

	if err := r.tx.Save(dto).Error; err != nil {
		return nil, core.NewError(core.SystemError, "タスクの更新に失敗しました。->"+err.Error())
	}
	result := dto.ToEntity()
	return result, nil
}

// 削除
func (r *taskRepository) DeleteById(id int) error {
	if err := r.tx.Delete(&task.Task{Id: id}).Error; err != nil {
		return core.NewError(core.SystemError, "タスクの削除に失敗しました。->"+err.Error())
	}
	return nil
}
