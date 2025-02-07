package language

import "gorm.io/gorm"

type LanguageRepositoryImpl struct {
	db *gorm.DB
}

func NewLanguageRepositoryImpl(db *gorm.DB) *LanguageRepositoryImpl {
	return &LanguageRepositoryImpl{db: db}
}
