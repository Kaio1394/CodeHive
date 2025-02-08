package code

import (
	"CodeHive/model"
	"CodeHive/repository/code"
	"context"
)

type CodeServiceImpl struct {
	rc *code.CodeRepositoryImpl
}

func NewCodeServiceImpl(rc *code.CodeRepositoryImpl) *CodeServiceImpl {
	return &CodeServiceImpl{rc: rc}
}

func (c *CodeServiceImpl) Create(ctx context.Context, code *model.Code) (int, error) {
	return c.rc.Create(ctx, code)
}
