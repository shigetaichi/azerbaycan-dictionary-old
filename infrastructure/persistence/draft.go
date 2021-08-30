package persistence

import (
	"context"
	"github.com/pkg/errors"
	"go-ddd/domain"
	"go-ddd/domain/entity"
	"go-ddd/domain/repository"
	"go-ddd/pkg/rdb"
	"go-ddd/pkg/util"
)

type draft struct{}

func NewDraft() repository.IDraft {
	return &draft{}
}

func (d draft) Create(ctx context.Context, draft *entity.Draft) (uint, error) {
	db := rdb.Get(ctx)

	if err := db.Create(draft).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return draft.ID, nil
}

func (d draft) GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Draft, uint, error) {
	db := rdb.Get(ctx)

	var res []*entity.Draft
	query := db.Model(&entity.Draft{}).Preload("User")
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	count, err := paging.GetCount(query, &entity.Draft{})
	if err != nil {
		return nil, 0, err
	}

	if err := query.Scopes(paging.Query()).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (d draft) GetById(ctx context.Context, id uint) (*entity.Draft, error) {
	db := rdb.Get(ctx)

	var res entity.Draft
	if err := db.
		Model(&entity.Draft{}).
		Where(&entity.Draft{
			HardDeleteModel: domain.HardDeleteModel{
				ID: id,
			},
		}).
		Last(&res).
		Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func (d draft) Update(ctx context.Context, word *entity.Draft) error {
	db := rdb.Get(ctx)

	if err := db.Model(&word).Updates(word).Error; err != nil {
		return err
	}
	return nil
}

func (d draft) Delete(ctx context.Context, id uint) error {
	db := rdb.Get(ctx)

	if err := db.
		Model(&entity.Draft{}).
		Delete(&entity.Draft{
			HardDeleteModel: domain.HardDeleteModel{
				ID: id,
			},
		}).
		Error; err != nil {
		return err
	}

	return nil
}
