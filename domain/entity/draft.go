package entity

import (
	"go-ddd/domain"
	"go-ddd/resource/request"
)

type Draft struct {
	domain.HardDeleteModel
	Name        string `json:"name"`
	Translation string `json:"translation"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`

	User *User `json:"user"`
}

func NewDraft(userId uint, dto *request.DraftCreate) *Draft {
	// validationの場合はentityで。
	return &Draft{
		UserID:      userId,
		Name:        dto.Name,
		Translation: dto.Translation,
		Description: dto.Description,
	}
}
