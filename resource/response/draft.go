package response

import "go-ddd/domain/entity"

type DraftGetAllResponse struct {
	Count  uint            `json:"count"`
	Drafts []*entity.Draft `json:"drafts"`
}
