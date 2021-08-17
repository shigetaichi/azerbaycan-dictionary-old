package repository

import (
	"context"
	"go-ddd/domain/entity"
	"go-ddd/pkg/util"
)

type IWord interface {
	Create(ctx context.Context, word *entity.Word) (uint, error)
	GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Word, uint, error)
	Update(ctx context.Context, word *entity.Word) error
}
