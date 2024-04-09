package task

import (
	"time"
)

type TaskEntity struct {
	id         int
	createdAt  time.Time
	updatedAt  time.Time
	userId     int
	title      string
	detail     string
	startDate  time.Time
	priorityId int
	categoryId int
	doneFlag   bool
}

// エンティティの作成
func NewEntity(userId int, title string, detail string, startDate time.Time, priorityId int, categoryId int) (*TaskEntity, error) {
	// バリデーション
	if err := validateTitle(title); err != nil {
		return nil, err
	}
	if err := validateStartDate(startDate); err != nil {
		return nil, err
	}
	return &TaskEntity{userId: userId, title: title, detail: detail, startDate: startDate, priorityId: priorityId, categoryId: categoryId, doneFlag: false}, nil
}

// タイトルの変更
func (e *TaskEntity) ChangeTitle(title string) error {
	// バリデーション
	if err := validateTitle(title); err != nil {
		return err
	}

	e.title = title
	return nil
}

// 詳細の変更
func (e *TaskEntity) ChangeDetail(detail string) error {
	e.detail = detail
	return nil
}

// 開始日の変更
func (e *TaskEntity) ChangeStartDate(startDate time.Time) error {
	// バリデーション
	if err := validateStartDate(startDate); err != nil {
		return err
	}

	e.startDate = startDate
	return nil
}

// 優先度の変更
func (e *TaskEntity) ChangePriority(priorityId int) error {
	e.priorityId = priorityId
	return nil
}

// カテゴリの変更
func (e *TaskEntity) ChangeCategory(categoryId int) error {
	e.categoryId = categoryId
	return nil
}

// タスクの完了
func (e *TaskEntity) Done() {
	e.doneFlag = true
}

// ゲッター　ID
func (e *TaskEntity) GetId() int {
	return e.id
}

// ゲッター　作成日時
func (e *TaskEntity) GetCreatedAt() time.Time {
	return e.createdAt
}

// ゲッター　更新日時
func (e *TaskEntity) GetUpdatedAt() time.Time {
	return e.updatedAt
}

// ゲッター　ユーザーID
func (e *TaskEntity) GetUserId() int {
	return e.userId
}

// ゲッター　タイトル
func (e *TaskEntity) GetTitle() string {
	return e.title
}

// ゲッター　詳細
func (e *TaskEntity) GetDetail() string {
	return e.detail
}

// ゲッター　開始日
func (e *TaskEntity) GetStartDate() time.Time {
	return e.startDate
}

// ゲッター　優先度ID
func (e *TaskEntity) GetProriotyId() int {
	return e.priorityId
}

// ゲッター　カテゴリID
func (e *TaskEntity) GetCategoryId() int {
	return e.categoryId
}

// ゲッター　完了フラグ
func (e *TaskEntity) GetDoneFlag() bool {
	return e.doneFlag
}
