package request

type DraftCreate struct {
	Name        string `json:"name" binding:"required"`
	Translation string `json:"translation" binding:"required"`
	Description string `json:"description"`
}

type DraftUpdate struct {
	Id          uint   `json:"id"`
	Name        string `json:"name" binding:"required"`
	Translation string `json:"translation" binding:"required"`
	Description string `json:"description"`
}
