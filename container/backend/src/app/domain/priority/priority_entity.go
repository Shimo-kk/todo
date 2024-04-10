package priority

import "time"

type PriorityEntity struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	name      string
}

// ゲッター　ID
func (e *PriorityEntity) GetId() int {
	return e.id
}

// ゲッター　作成日時
func (e *PriorityEntity) GetCreatedAt() time.Time {
	return e.createdAt
}

// ゲッター　更新日時
func (e *PriorityEntity) GetUpdatedAt() time.Time {
	return e.updatedAt
}

// ゲッター　名称
func (e *PriorityEntity) GetName() string {
	return e.name
}
