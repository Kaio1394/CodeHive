package language

import (
	"CodeHive/model"
	"context"
)

type LanguageRepository interface {
	CreateLanguage(ctx context.Context, language *model.Language) error
	GetLanguageById(ctx context.Context, id int) (model.Language, error)
	UpdateLanguage(ctx context.Context, id int, languageModel model.Language) error
}
