package controller

import "CodeHive/service/code"

type CodeController struct {
	sc *code.CodeServiceImpl
}

func NewCodeController(sc *code.CodeServiceImpl) *CodeController {
	return &CodeController{sc}
}
