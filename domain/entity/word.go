package entity

import (
	"go-ddd/domain"
	"go-ddd/resource/request"
)

type Word struct {
	domain.SoftDeleteModel
	Name        string `json:"name"`
	Translation string `json:"translation"`
	Star        int    `json:"star"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`

	User *User `json:"user"`

	Users []*User `gorm:"many2many:bookmarks;"`
}

func NewWord(userId uint, dto *request.WordCreate) *Word {
	// validationの場合はentityで。
	return &Word{
		UserID:      userId,
		Name:        dto.Name,
		Translation: dto.Translation,
		Description: dto.Description,
		Star:        0,
	}
}
