package request

type WordCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
