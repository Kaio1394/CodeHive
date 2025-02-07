package code

import (
	"CodeHive/model"
	"context"
)

type CodeService interface {
	Create(ctx context.Context, code *model.Code) (int, error)
}
