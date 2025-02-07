package code

import (
	"CodeHive/logger"
	"CodeHive/model"
	"context"
	"gorm.io/gorm"
)

type CodeRepositoryImpl struct {
	db *gorm.DB
}

func (c CodeRepositoryImpl) Create(ctx context.Context, code *model.Code) (int, error) {
	if err := c.db.WithContext(ctx).Create(&code).Error; err != nil {
		logger.Log.Error("Error to insert a new server: " + err.Error())
		return 0, err
	}
	return code.Id, nil
}

func NewCodeRepositoryImpl(db *gorm.DB) *CodeRepositoryImpl {
	return &CodeRepositoryImpl{db: db}
}
