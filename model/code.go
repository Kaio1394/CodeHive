package model

import "time"

type Code struct {
	Id         int       `json:"id" gorm:"primaryKey"`
	Tag        string    `json:"tag" gorm:"not null"`
	Script     string    `json:"script" gorm:"not null"`
	LanguageID int       `json:"language_id" gorm:"not null;index"`
	Language   Language  `json:"language" gorm:"foreignKey:LanguageID;references:Id"`
	CreateAt   time.Time `json:"createAt" gorm:"column:create_at;default:CURRENT_TIMESTAMP"`
}

func (Code) TableName() string {
	return "t_register_codes"
}
