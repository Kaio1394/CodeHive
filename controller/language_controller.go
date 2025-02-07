package controller

import "CodeHive/service/language"

type LanguageController struct {
	ls *language.LanguageServiceImpl
}

func NewLanguageController(ls *language.LanguageServiceImpl) *LanguageController {
	return &LanguageController{ls: ls}
}
