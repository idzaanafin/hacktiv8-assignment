package models

type Product struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SourceID    string  `json:"source_id"`
}

type Source struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Transaction struct {
	ID        string  `json:"id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Total     float64 `json:"total"`
}
