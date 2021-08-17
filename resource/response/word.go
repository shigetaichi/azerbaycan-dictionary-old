package response

import "go-ddd/domain/entity"

type WordGetAllResponse struct {
	Count uint           `json:"count"`
	Words []*entity.Word `json:"words"`
}
