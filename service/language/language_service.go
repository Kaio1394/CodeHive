package language

import (
	"CodeHive/model"
	"context"
)

type LanguageService interface {
	CreateLanguage(ctx context.Context, language *model.Language) error
	GetLanguageById(ctx context.Context, id int) (model.Language, error)
	UpdateLanguage(ctx context.Context, id int, language model.Language) error
}
