package usecase

import (
	"context"
	"go-ddd/domain/entity"
	"go-ddd/domain/repository"
	"go-ddd/pkg/util"
	"go-ddd/resource/request"
)

type IWord interface {
	Create(ctx context.Context, userId uint, req *request.WordCreate) (uint, error)
	GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Word, uint, error)
	GetById(ctx context.Context, id uint) (*entity.Word, error)
	Update(ctx context.Context, word *entity.Word) error
}
type word struct {
	wordRepo repository.IWord
}

func NewWord(wr repository.IWord) IWord {
	return &word{
		wordRepo: wr,
	}
}

func (w word) Create(ctx context.Context, userId uint, req *request.WordCreate) (uint, error) {
	word := entity.NewWord(userId, req)
	id, err := w.wordRepo.Create(ctx, word)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (w word) GetAll(ctx context.Context, keyword string, paging *util.Paging) ([]*entity.Word, uint, error) {
	res, count, err := w.wordRepo.GetAll(ctx, keyword, paging)
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}

func (w word) GetById(ctx context.Context, id uint) (*entity.Word, error) {
	res, err := w.wordRepo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (w word) Update(ctx context.Context, word *entity.Word) error {
	if err := w.wordRepo.Update(ctx, word); err != nil {
		return err
	}
	return nil
}
