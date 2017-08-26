package types

type Item struct {
	Id int                          `json:"id"`
	Name string                     `json:"name"`
	Description string              `json:"description"`
	Attributes map[string] string   `json:"attributes"`
	Tags []string                   `json:"tags"`
	Price float32                   `json:"price"`
	WholesalePrice float32          `json:"wholesalePrice"`
	Quantity float32                `json:"quantity"`
}