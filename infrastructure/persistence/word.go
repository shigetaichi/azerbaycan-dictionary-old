package persistence

import (
	"context"
	"github.com/pkg/errors"
	"go-ddd/domain/entity"
	"go-ddd/domain/repository"
	"go-ddd/pkg/rdb"
	"go-ddd/pkg/util"
)

type word struct{}

func NewWord() repository.IWord {
	return &word{}
}

func (w word) Create(ctx context.Context, word *entity.Word) (uint, error) {
	db := rdb.Get(ctx)

	if err := db.Create(word).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return word.ID, nil
}

func (w word) GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Word, uint, error) {
	db := rdb.Get(ctx)

	var res []*entity.Word
	query := db.Model(&entity.Word{}).Preload("User")
	if keyword != "" {
		query = query.Where("name LIKE ?", "%"+keyword+"%")
	}

	count, err := paging.GetCount(query, &entity.Word{})
	if err != nil {
		return nil, 0, err
	}

	if err := query.Scopes(paging.Query()).Find(&res).Error; err != nil {
		return nil, 0, err
	}
	return res, count, nil
}
