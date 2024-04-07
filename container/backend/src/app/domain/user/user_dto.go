package user

import "time"

type User struct {
	Id        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Email     string
	Password  string
}

// エンティティをDTOに変換
func ToDtoFromEntity(e *UserEntity) *User {
	return &User{
		Id:        e.id,
		CreatedAt: e.createdAt,
		UpdatedAt: e.updatedAt,
		Name:      e.name,
		Email:     e.email,
		Password:  e.password,
	}
}

// DTOをエンティティに変換
func (d *User) ToEntity() *UserEntity {
	return &UserEntity{
		id:        d.Id,
		createdAt: d.CreatedAt,
		updatedAt: d.UpdatedAt,
		name:      d.Name,
		email:     d.Email,
		password:  d.Password,
	}
}
