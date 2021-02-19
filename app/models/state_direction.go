package models

type StateDirection struct {
	StateClassId     int       `json:"state_class_id"`
	StateDirectionId int       `json:"state_direction_id"`
	StateFrom        string       `json:"state_from"`
	StateTo          string       `json:"state_to"`
	IsDeleted        int       `json:"is_deleted"`
	DeletedAt        DateTime `json:"deleted_at"`
	CreatedAt        DateTime `json:"created_at"`
	UpdatedAt        DateTime `json:"updated_at"`
}
