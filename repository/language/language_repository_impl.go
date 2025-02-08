package language

import (
	"CodeHive/logger"
	"CodeHive/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type LanguageRepositoryImpl struct {
	db *gorm.DB
}

func NewLanguageRepositoryImpl(db *gorm.DB) *LanguageRepositoryImpl {
	return &LanguageRepositoryImpl{db: db}
}

func (lr *LanguageRepositoryImpl) CreateLanguage(ctx context.Context, language *model.Language) error {

	err := lr.db.WithContext(ctx).Create(&language)
	if err != nil {
		logger.Log.Error("error creating language", "error", err)
	}
	return nil
}

func (lr *LanguageRepositoryImpl) GetLanguageById(ctx context.Context, id int) (model.Language, error) {
	var language model.Language
	result := lr.db.WithContext(ctx).First(&language, id) // Alterado de Find para First
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return model.Language{}, errors.New("language not found")
		}
		return model.Language{}, result.Error
	}
	return language, nil
}

func (lr *LanguageRepositoryImpl) UpdateLanguage(ctx context.Context, id int, languageModel model.Language) error {
	var language model.Language
	result := lr.db.WithContext(ctx).First(&language, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return errors.New("language not found")
		}
		return result.Error
	}
	result = lr.db.WithContext(ctx).Model(&language).Updates(languageModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
