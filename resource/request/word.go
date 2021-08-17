package request

type WordCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type WordUpdate struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Star        int    `json:"star"`
}

type WordUpdateStar struct {
	Id   uint `json:"id"`
	Star int  `json:"star"`
}
