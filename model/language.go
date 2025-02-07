package model

type Language struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"not null"`
	Extension string `json:"extension" gorm:"not null"`
}

func (Language) TableName() string {
	return "t_register_language_programming"
}
