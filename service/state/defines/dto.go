package defines

import "owen2020/app/models"

type StateDTO struct {
	models.State
	DbName        string    `form:"db_name" json:"db_name" xorm:"default '' VARCHAR(255)"`
	TableName     string    `form:"table_name" json:"table_name" xorm:"default '' VARCHAR(255)"`
	FieldName     string    `form:"field_name" json:"field_name" xorm:"default '' VARCHAR(255)"`
	StateClassName string `form:"state_class_name" json:"state_class_name"`
}

type StateDirectionDto struct {
	models.StateDirection
	StateFromDesc string `form:"state_from_desc" json:"state_from_desc"`
	StateToDesc   string `form:"state_to_desc" json:"state_to_desc"`
}
