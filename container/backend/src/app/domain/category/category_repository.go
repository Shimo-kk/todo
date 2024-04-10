package category

type ICategoryRepository interface {
	Insert(entity *CategoryEntity) (*CategoryEntity, error)
	FindAll(userId int) (*[]CategoryEntity, error)
	FindById(id int) (*CategoryEntity, error)
	Update(entity *CategoryEntity) (*CategoryEntity, error)
	DeleteById(id int) error
}
