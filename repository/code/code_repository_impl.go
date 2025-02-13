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

func (c *CodeRepositoryImpl) Create(ctx context.Context, code *model.Code) (int, error) {
	if err := c.db.WithContext(ctx).Create(&code).Error; err != nil {
		logger.Log.Error("Error to insert a new server: " + err.Error())
		return 0, err
	}
	return code.Id, nil
}

func (c *CodeRepositoryImpl) GetCodeById(ctx context.Context, id int) (model.Code, error) {
	var code model.Code
	result := c.db.WithContext(ctx).Find(&code, &id)
	if result.Error != nil {
		logger.Log.Error("Error to get a code: " + result.Error.Error())
		return model.Code{}, result.Error
	}
	return code, nil
}
func (c *CodeRepositoryImpl) GetListCode(ctx context.Context) ([]model.Code, error) {
	var codes []model.Code
	if err := c.db.WithContext(ctx).Find(&codes).Error; err != nil {
		logger.Log.Error("Error to get a list of code: " + err.Error())
		return nil, err
	}
	return codes, nil
}
func (c *CodeRepositoryImpl) Delete(ctx context.Context, id int) error {
	var cod model.Code
	if err := c.db.WithContext(ctx).First(&cod, id).Error; err != nil {
		logger.Log.Error("Error to get a code: " + err.Error())
		return err
	}
	if err := c.db.WithContext(ctx).Delete(&cod).Error; err != nil {
		logger.Log.Error("Error to delete a code: " + err.Error())
		return err
	}
	return nil
}

func NewCodeRepositoryImpl(db *gorm.DB) *CodeRepositoryImpl {
	return &CodeRepositoryImpl{db: db}
}
