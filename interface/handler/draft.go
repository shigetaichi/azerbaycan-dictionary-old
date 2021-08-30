package handler

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/ken109/gin-jwt"
	"go-ddd/constant"
	"go-ddd/domain"
	"go-ddd/domain/entity"
	"go-ddd/pkg/util"
	"go-ddd/pkg/xerrors"
	"go-ddd/resource/request"
	"go-ddd/resource/response"
	"go-ddd/usecase"
	"net/http"
	"strconv"
)

type Draft struct {
	draftUseCase usecase.IDraft
}

func NewDraft(route *gin.RouterGroup, duc usecase.IDraft) {
	handler := Draft{
		draftUseCase: duc,
	}

	route.Use(jwt.Verify(constant.DefaultRealm))
	get(route, "", handler.GetAll)
	get(route, "id/:id", handler.GetById)
	post(route, "", handler.Create)
	put(route, "", handler.Update)
	patch(route, "id/:id", handler.Publish)
}

func (d Draft) Create(c *gin.Context) error {
	var req request.DraftCreate

	if !bind(c, &req) {
		return nil
	}

	id := jwt.GetClaims(c)["id"].(float64)
	res, err := d.draftUseCase.Create(newCtx(), uint(id), &req)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, res)
	return nil
}

func (d Draft) GetAll(c *gin.Context) error {
	paging := util.NewPaging(c)
	res, count, err := d.draftUseCase.GetAll(newCtx(), c.Query("keyword"), paging)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, response.DraftGetAllResponse{
		Count:  count,
		Drafts: res,
	})
	return nil
}
func (d Draft) GetById(c *gin.Context) error {
	wid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return xerrors.NewExpected(http.StatusNotFound, "Invalid Draft Id")
	}
	res, err := d.draftUseCase.GetById(newCtx(), uint(wid))
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, res)
	return nil
}

func (d Draft) Update(c *gin.Context) error {
	var req request.DraftUpdate
	if !bind(c, &req) {
		return nil
	}

	if err := d.draftUseCase.Update(newCtx(), &entity.Draft{
		HardDeleteModel: domain.HardDeleteModel{
			ID: req.Id,
		},
		Name:        req.Name,
		Translation: req.Translation,
		Description: req.Description,
	}); err != nil {
		return err
	}

	c.Status(http.StatusOK)
	return nil
}

func (d Draft) Publish(c *gin.Context) error {
	wid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return xerrors.NewExpected(http.StatusNotFound, "Invalid Draft Id")
	}

	res, err := d.draftUseCase.Publish(newCtx(), uint(wid))
	if err != nil {
		return err
	}
	c.JSON(http.StatusCreated, res)
	return nil
}
