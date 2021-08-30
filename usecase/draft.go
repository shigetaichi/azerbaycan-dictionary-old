package usecase

import (
	"context"
	"go-ddd/domain/entity"
	"go-ddd/domain/repository"
	"go-ddd/pkg/rdb"
	"go-ddd/pkg/util"
	"go-ddd/resource/request"
)

type IDraft interface {
	Create(ctx context.Context, userId uint, req *request.DraftCreate) (uint, error)
	GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Draft, uint, error)
	GetById(ctx context.Context, id uint) (*entity.Draft, error)
	Update(ctx context.Context, word *entity.Draft) error
	Publish(ctx context.Context, id uint) (uint, error)
}
type draft struct {
	draftRepo repository.IDraft
	wordRepo  repository.IWord
}

func NewDraft(
	dr repository.IDraft,
	wr repository.IWord,
) IDraft {
	return &draft{
		draftRepo: dr,
		wordRepo:  wr,
	}
}

func (d draft) Create(ctx context.Context, userId uint, req *request.DraftCreate) (uint, error) {
	draft := entity.NewDraft(userId, req)
	id, err := d.draftRepo.Create(ctx, draft)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (d draft) GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Draft, uint, error) {
	res, count, err := d.draftRepo.GetAll(ctx, keyword, paging)
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (d draft) GetById(ctx context.Context, id uint) (*entity.Draft, error) {
	res, err := d.draftRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d draft) Update(ctx context.Context, word *entity.Draft) error {
	if err := d.draftRepo.Update(ctx, word); err != nil {
		return err
	}
	return nil
}

func (d draft) Publish(ctx context.Context, id uint) (uint, error) {
	res, err := d.draftRepo.GetById(ctx, id)
	if err != nil {
		return 0, err
	}

	word := entity.NewWord(res.UserID, &request.WordCreate{
		Name:        res.Name,
		Translation: res.Translation,
		Description: res.Description,
	})

	var wid uint
	err = rdb.Transaction(ctx, func(ctx context.Context) error {
		wid, err = d.wordRepo.Create(ctx, word)
		if err != nil {
			return err
		}

		err = d.draftRepo.Delete(ctx, id)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return wid, nil
}
