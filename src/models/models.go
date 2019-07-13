// data structures for the project 

package models

type User struct {
	ID 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	Email		string 	`json:"email`
	Password	string 	`json:"password"`
	Balance		int 	`json:"balance"`
}

type Transaction struct {

	ID			int 	`json:"id"`
	Type 		string 	`json:"type"`
	Symbol 		string 	`json:"symbol"`
	Quantity 	int 	`json:"quantity"`
	Price 		int 	`json:"price"`
}

type Holding struct {

	ID			int 	`json:"id"`
	Symbol 		string 	`json:"symbol"`
	Quantity 	int 	`json:"quantity"`
}

type JWT struct {
	Token string 	`json:"token"`
}

type Error struct {
	Message string `json:"message"`
}
