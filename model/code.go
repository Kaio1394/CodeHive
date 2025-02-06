package model

import "time"

type Code struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Tag      string `json:"tag"`
	Script   string
	CreateAt time.Time `json:"createAt" gorm:"column:create_at;default:CURRENT_TIMESTAMP"`
}

func (Code) TableName() string {
	return "t_register_codes"
}
