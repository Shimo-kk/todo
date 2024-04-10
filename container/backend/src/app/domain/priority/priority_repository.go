package priority

type IPriorityRepository interface {
	FindAll() (*[]PriorityEntity, error)
}
