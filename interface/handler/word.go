package handler

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/ken109/gin-jwt"
	"go-ddd/domain"
	"go-ddd/domain/entity"
	"go-ddd/pkg/util"
	"go-ddd/resource/request"
	"go-ddd/resource/response"
	"go-ddd/usecase"
	"net/http"
)

type Word struct {
	wordUseCase usecase.IWord
}

func NewWord(route *gin.RouterGroup, wuc usecase.IWord) {
	handler := Word{
		wordUseCase: wuc,
	}

	get(route, "", handler.GetAll)
	post(route, "", handler.Create)
	put(route, "", handler.Update)
	patch(route, "", handler.UpdateStar)
}

func (w Word) Create(c *gin.Context) error {
	var req request.WordCreate

	if !bind(c, &req) {
		return nil
	}

	id := jwt.GetClaims(c)["id"].(float64)
	res, err := w.wordUseCase.Create(newCtx(), uint(id), &req)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, res)
	return nil
}

func (w Word) GetAll(c *gin.Context) error {
	paging := util.NewPaging(c)
	res, count, err := w.wordUseCase.GetAll(newCtx(), c.Query("keyword"), paging)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, response.WordGetAllResponse{
		Count: count,
		Words: res,
	})
	return nil
}

func (w Word) Update(c *gin.Context) error {
	var req request.WordUpdate
	if !bind(c, &req) {
		return nil
	}

	if err := w.wordUseCase.Update(newCtx(), &entity.Word{
		SoftDeleteModel: domain.SoftDeleteModel{
			ID: req.Id,
		},
		Name:        req.Name,
		Description: req.Description,
		Star:        req.Star,
	}); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (w Word) UpdateStar(c *gin.Context) error {
	var req request.WordUpdateStar
	if !bind(c, &req) {
		return nil
	}

	if err := w.wordUseCase.Update(newCtx(), &entity.Word{
		SoftDeleteModel: domain.SoftDeleteModel{
			ID: req.Id,
		},
		Star: req.Star,
	}); err != nil {
		return err
	}
	c.Status(http.StatusOK)
	return nil
}
