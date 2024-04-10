package priority

import "time"

type Priority struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

// エンティティをDTOに変換
func ToDtoFromEntity(e *PriorityEntity) *Priority {
	return &Priority{
		Id:        e.id,
		CreatedAt: e.createdAt,
		UpdatedAt: e.updatedAt,
		Name:      e.name,
	}
}

// DTOをエンティティに変換
func (d *Priority) ToEntity() *PriorityEntity {
	return &PriorityEntity{
		id:        d.Id,
		createdAt: d.CreatedAt,
		updatedAt: d.UpdatedAt,
		name:      d.Name,
	}
}
