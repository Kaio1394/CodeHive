package language

import (
	"CodeHive/model"
	"CodeHive/repository/language"
	"context"
)

type LanguageServiceImpl struct {
	lr *language.LanguageRepositoryImpl
}

func NewLanguageServiceImpl(lr *language.LanguageRepositoryImpl) *LanguageServiceImpl {
	return &LanguageServiceImpl{lr: lr}
}

func (ls *LanguageServiceImpl) CreateLanguage(ctx context.Context, language *model.Language) error {
	return ls.lr.CreateLanguage(ctx, language)
}

func (ls *LanguageServiceImpl) GetLanguageById(ctx context.Context, id int) (model.Language, error) {
	return ls.lr.GetLanguageById(ctx, id)
}

func (ls *LanguageServiceImpl) UpdateLanguage(ctx context.Context, id int, language model.Language) error {
	return ls.lr.UpdateLanguage(ctx, id, language)
}

func (ls *LanguageServiceImpl) DeleteLanguageById(ctx context.Context, id int) error {
	return ls.lr.DeleteLanguageById(ctx, id)
}

func (ls *LanguageServiceImpl) GetAllLanguages(ctx context.Context) ([]model.Language, error) {
	return ls.lr.GetListLanguage(ctx)
}
