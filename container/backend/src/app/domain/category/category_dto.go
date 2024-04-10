package category

import "time"

type Category struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// エンティティをDTOに変換
func ToDtoFromEntity(e *CategoryEntity) *Category {
	return &Category{
		Id:        e.id,
		CreatedAt: e.createdAt,
		UpdatedAt: e.updatedAt,
		Name:      e.name,
	}
}

// DTOをエンティティに変換
func (d *Category) ToEntity() *CategoryEntity {
	return &CategoryEntity{
		id:        d.Id,
		createdAt: d.CreatedAt,
		updatedAt: d.UpdatedAt,
		name:      d.Name,
	}
}
