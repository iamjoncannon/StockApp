package models

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email`
	Password string  `json:"password"`
	Balance  float32 `json:"balance"`
}

type Transaction struct {
	ID       int     `json:"id"`
	Type     string  `json:"type"`
	Symbol   string  `json:"symbol"`
	Quantity float32 `json:"quantity"`
	Price    float32 `json:"price"`
}

type Holding struct {
	ID       int     `json:"id"`
	Symbol   string  `json:"symbol"`
	Quantity float32 `json:"quantity"`
}

type Error struct {
	Message string `json:"message"`
}
