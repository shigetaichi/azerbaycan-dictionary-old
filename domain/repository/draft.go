package repository

import (
	"context"
	"go-ddd/domain/entity"
	"go-ddd/pkg/util"
)

type IDraft interface {
	Create(ctx context.Context, word *entity.Draft) (uint, error)
	GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Draft, uint, error)
	GetById(ctx context.Context, id uint) (*entity.Draft, error)
	Update(ctx context.Context, word *entity.Draft) error
	Delete(ctx context.Context, id uint) error
}
