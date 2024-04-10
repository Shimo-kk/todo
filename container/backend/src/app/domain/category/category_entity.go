package category

import "time"

type CategoryEntity struct {
	id        int
	createdAt time.Time
	updatedAt time.Time
	name      string
}

// エンティティの作成
func NewEntity(name string) (*CategoryEntity, error) {
	// バリデーション
	if err := validateName(name); err != nil {
		return nil, err
	}

	return &CategoryEntity{name: name}, nil
}

// 名称の変更
func (e *CategoryEntity) ChangeName(name string) error {
	// バリデーション
	if err := validateName(name); err != nil {
		return err
	}

	e.name = name
	return nil
}

// ゲッター　ID
func (e *CategoryEntity) GetId() int {
	return e.id
}

// ゲッター　作成日時
func (e *CategoryEntity) GetCreatedAt() time.Time {
	return e.createdAt
}

// ゲッター　更新日時
func (e *CategoryEntity) GetUpdatedAt() time.Time {
	return e.updatedAt
}

// ゲッター　名称
func (e *CategoryEntity) GetName() string {
	return e.name
}
