package usecase

import (
	"todo/app/application/interface/database"
	"todo/app/application/interface/usecase"
	"todo/app/application/schema"
	"todo/app/core"
	"todo/app/domain/task"
)

type taskUsecase struct {
	databaseHandller database.IDatabaseHandller
}

// タスクユースケースの作成
func NewTaskUsecase(databaseHandller database.IDatabaseHandller) usecase.ITaskUsecase {
	return &taskUsecase{databaseHandller: databaseHandller}
}

// タスクの作成
func (u *taskUsecase) CreateTask(userId int, data schema.TaskCreateModel) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		taskRepository := rf.GetTaskRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// タスクのエンティティを作成
		taskEntity, err := task.NewEntity(userEntity.GetId(), data.Title, data.Detail, data.StartDate, data.PriorityId, data.CategoryId)
		if err != nil {
			return err
		}

		// タスクを挿入
		_, err = taskRepository.Insert(taskEntity)
		if err != nil {
			return err
		}

		return nil
	})
}

// タスクの更新
func (u *taskUsecase) UpdateTask(userId int, data schema.TaskUpdateModel) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		taskRepository := rf.GetTaskRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// タスクを取得
		taskEntity, err := taskRepository.FindById(data.Id)
		if err != nil {
			return err
		}
		if taskEntity == nil {
			return core.NewError(core.NotFoundError, "タスクが存在しません。")
		}

		// 各種プロパティの変更
		if err := taskEntity.ChangeTitle(data.Title); err != nil {
			return err
		}
		if err := taskEntity.ChangeDetail(data.Detail); err != nil {
			return err
		}
		if err := taskEntity.ChangeStartDate(data.StartDate); err != nil {
			return err
		}
		if err := taskEntity.ChangePriority(data.PriorityId); err != nil {
			return err
		}
		if err := taskEntity.ChangeCategory(data.CategoryId); err != nil {
			return err
		}

		// タスクの更新
		_, err = taskRepository.Update(taskEntity)
		if err != nil {
			return err
		}

		return nil
	})
}

// タスクの取得
func (u *taskUsecase) GetTask(userId int, id int) (*schema.TaskReadModel, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	userRepository := repositoryFactory.GetUserRepository()
	taskRepository := repositoryFactory.GetTaskRepository()

	// ユーザーを取得
	userEntity, err := userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
	}

	// タスクを取得
	taskEntity, err := taskRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	if taskEntity == nil {
		return nil, core.NewError(core.NotFoundError, "タスクが存在しません。")
	}

	// スキーマへ変換
	result := schema.TaskReadModel{
		Id:         taskEntity.GetId(),
		Title:      taskEntity.GetTitle(),
		Detail:     taskEntity.GetDetail(),
		StartDate:  taskEntity.GetStartDate(),
		PriorityId: taskEntity.GetProriotyId(),
		CategoryId: taskEntity.GetCategoryId(),
		DoneFlag:   taskEntity.GetDoneFlag(),
	}

	return &result, nil
}

// タスクの削除
func (u *taskUsecase) DeleteTask(userId int, id int) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		taskRepository := rf.GetTaskRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// タスクの削除
		if err := taskRepository.DeleteById(id); err != nil {
			return err
		}

		return nil
	})
}

// タスクの完了
func (u *taskUsecase) DoneTask(userId int, id int) error {
	return u.databaseHandller.Transaction(func(rf database.IRepositoryFactory) error {
		userRepository := rf.GetUserRepository()
		taskRepository := rf.GetTaskRepository()

		// ユーザーを取得
		userEntity, err := userRepository.FindById(userId)
		if err != nil {
			return err
		}
		if userEntity == nil {
			return core.NewError(core.NotFoundError, "ユーザーが存在しません。")
		}

		// タスクを取得
		taskEntity, err := taskRepository.FindById(id)
		if err != nil {
			return err
		}
		if taskEntity == nil {
			return core.NewError(core.NotFoundError, "タスクが存在しません。")
		}

		// タスクを完了
		taskEntity.Done()

		// タスクの更新
		_, err = taskRepository.Update(taskEntity)
		if err != nil {
			return err
		}

		return nil
	})
}

// タスクの取得
func (u *taskUsecase) GetAllTask(userId int) (*[]schema.TaskReadModel, error) {
	repositoryFactory := u.databaseHandller.GetRepositoryFactory()
	userRepository := repositoryFactory.GetTaskRepository()
	taskRepository := repositoryFactory.GetTaskRepository()

	// ユーザーを取得
	userEntity, err := userRepository.FindById(userId)
	if err != nil {
		return nil, err
	}
	if userEntity == nil {
		return nil, core.NewError(core.NotFoundError, "ユーザーが存在しません。")
	}

	// タスクを全件取得
	taskEntityList, err := taskRepository.FindAll(userEntity.GetId())
	if err != nil {
		return nil, err
	}

	// スキーマへ変換
	result := []schema.TaskReadModel{}
	for _, entity := range *taskEntityList {
		model := schema.TaskReadModel{
			Id:         entity.GetId(),
			Title:      entity.GetTitle(),
			Detail:     entity.GetDetail(),
			StartDate:  entity.GetStartDate(),
			PriorityId: entity.GetProriotyId(),
			CategoryId: entity.GetCategoryId(),
			DoneFlag:   entity.GetDoneFlag(),
		}
		result = append(result, model)
	}

	return &result, nil
}
