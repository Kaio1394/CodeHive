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

func (c *CodeServiceImpl) GetListCode(ctx context.Context) ([]model.Code, error) {
	return c.rc.GetListCode(ctx)
}

func (c *CodeServiceImpl) Delete(ctx context.Context, id int) error {
	return c.rc.Delete(ctx, id)
}

func (c *CodeServiceImpl) GetCodeById(ctx context.Context, id int) (model.Code, error) {
	return c.rc.GetCodeById(ctx, id)
}
