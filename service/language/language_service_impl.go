package language

import "CodeHive/repository/language"

type LanguageServiceImpl struct {
	lr *language.LanguageRepositoryImpl
}

func NewLanguageServiceImpl(lr *language.LanguageRepositoryImpl) *LanguageServiceImpl {
	return &LanguageServiceImpl{lr: lr}
}
