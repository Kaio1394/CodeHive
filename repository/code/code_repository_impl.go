package code

import "gorm.io/gorm"

type CodeRepositoryImpl struct {
	db *gorm.DB
}
