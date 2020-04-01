package entities

type Product struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Price    float32   `json:"price"`
	Quantity int       `json:"quantity"`
	Status   bool      `json:"status"`
	Category Category  `json:"category"`
	Comments []Comment `json:"comments"`
}
