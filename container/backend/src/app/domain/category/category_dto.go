package category

import "time"

type Category struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserId    int
	Name      string
}

// エンティティをDTOに変換
func ToDtoFromEntity(e *CategoryEntity) *Category {
	return &Category{
		Id:        e.id,
		CreatedAt: e.createdAt,
		UpdatedAt: e.updatedAt,
		UserId:    e.userId,
		Name:      e.name,
	}
}

// DTOをエンティティに変換
func (d *Category) ToEntity() *CategoryEntity {
	return &CategoryEntity{
		id:        d.Id,
		createdAt: d.CreatedAt,
		updatedAt: d.UpdatedAt,
		userId:    d.UserId,
		name:      d.Name,
	}
}
