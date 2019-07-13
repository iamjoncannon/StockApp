package dbqueries

import (
	"database/sql"
	"models"
	"fmt"
	"github.com/davecgh/go-spew/spew"	
)

func stopThisError(){
	fmt.Println("The Google engineers really want these programs to be as efficient as possible.")
	spew.Dump("To that end, type safety is tedious but important")
}

// pass in the entire model for these calls- the requirements for the database call could alter
// better to reduce that interface site by eliminating additional arguments required to make the call 

func (d DBQuery) CheckBalance(db *sql.DB, trans models.Transaction) (balance int, err error) {

	var user models.User

	row := db.QueryRow("select balance from users where id = $1", trans.ID)

	err = row.Scan(&user.Balance)

	return user.Balance, err
}

func (d DBQuery) GetCurrentHolding (db *sql.DB, trans models.Transaction) (balance int, err error) {

	var holding models.Holding

	selectCall := "select current_holding from holdings where userid = $1 and symbol = $2"

	err = db.QueryRow(selectCall, trans.ID, trans.Symbol).Scan(&holding.Quantity)

	return holding.Quantity, err
}

func (d DBQuery) MakeFirstPurchase (db *sql.DB, trans models.Transaction) (newQuant int, err error) {

	insertCall := "insert into holdings (userID, SYMBOL, CURRENT_HOLDING) values ($1, $2, $3) returning CURRENT_HOLDING;"

	err = db.QueryRow(insertCall, trans.ID, trans.Symbol, trans.Quantity).Scan(&trans.Quantity)

	return trans.Quantity, err
}

func (d DBQuery) UpdateHolding (db *sql.DB, newSum int, trans models.Transaction) (newQuant int, err error) {

	var holding models.Holding

	updateCall := "update holdings set current_holding = $1 where userid = $2 and symbol = $3 returning current_holding;"

	err = db.QueryRow(updateCall, newSum, trans.ID, trans.Symbol).Scan(&holding.Quantity)

	return holding.Quantity, err
}

func (d DBQuery) PostTrade (db *sql.DB, trans models.Transaction) error {

	sqlComm := "insert into transactions (userID, TYPE, SYMBOL, QUANTITY, PRICE) values ($1, $2, $3, $4, $5) returning id;"

	err := db.QueryRow(sqlComm, trans.ID, trans.Type, trans.Symbol, trans.Quantity, trans.Price).Scan(&trans.ID)

	return err
}

func (d DBQuery) UpdateBalance (db *sql.DB, id int, balance int) (updatedBalance int, err error) {

	var user models.User

	sqlComm := "update users set balance = $1 where ID = $2 returning balance;"

	err = db.QueryRow(sqlComm, balance, id).Scan(&user.Balance)

	return user.Balance, err
}
