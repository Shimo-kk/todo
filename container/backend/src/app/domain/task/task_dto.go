package task

import "time"

type Task struct {
	Id         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserId     int
	Title      string
	Detail     string
	StartDate  time.Time
	PriorityId int
	CategoryId int
	DoneFlag   bool
}

// エンティティをDTOに変換
func ToDtoFromEntity(ue *TaskEntity) *Task {
	return &Task{
		Id:         ue.id,
		CreatedAt:  ue.createdAt,
		UpdatedAt:  ue.updatedAt,
		UserId:     ue.userId,
		Title:      ue.title,
		Detail:     ue.detail,
		StartDate:  ue.startDate,
		PriorityId: ue.priorityId,
		CategoryId: ue.categoryId,
		DoneFlag:   ue.doneFlag,
	}
}

// DTOをエンティティに変換
func (ud *Task) ToEntity() *TaskEntity {
	return &TaskEntity{
		id:         ud.Id,
		createdAt:  ud.CreatedAt,
		updatedAt:  ud.UpdatedAt,
		userId:     ud.UserId,
		title:      ud.Title,
		detail:     ud.Detail,
		startDate:  ud.StartDate,
		priorityId: ud.PriorityId,
		categoryId: ud.CategoryId,
		doneFlag:   ud.DoneFlag,
	}
}
