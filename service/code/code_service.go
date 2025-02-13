package code

import (
	"CodeHive/model"
	"context"
)

type CodeService interface {
	Create(ctx context.Context, code *model.Code) (int, error)
	Delete(ctx context.Context, id int) error
	GetListCode(ctx context.Context) ([]model.Code, error)
	GetCodeById(ctx context.Context, id int) (model.Code, error)
}
