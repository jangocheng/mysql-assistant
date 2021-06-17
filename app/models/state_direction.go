package models

type StateDirection struct {
	StateClassId     int      `form:"state_class_id" json:"state_class_id"`
	StateDirectionId int      `form:"state_direction_id" json:"state_direction_id"`
	StateFrom        string   `form:"state_from" json:"state_from"`
	StateTo          string   `form:"state_to" json:"state_to"`
	Label            string   `form:"label" json:"label"`
	IsDeleted        int      `form:"is_deleted" json:"is_deleted"`
	DeletedAt        DateTime `form:"deleted_at" json:"deleted_at"`
	CreatedAt        DateTime `form:"created_at" json:"created_at"`
	UpdatedAt        DateTime `form:"updated_at" json:"updated_at"`
}
