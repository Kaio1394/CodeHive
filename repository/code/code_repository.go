package code

import (
	"CodeHive/model"
	"context"
)

type CodeRepository interface {
	Create(ctx context.Context, code *model.Code) (int, error)
	GetCodeById(ctx context.Context, id int) (model.Code, error)
}
