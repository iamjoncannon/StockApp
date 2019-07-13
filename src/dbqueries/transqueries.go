package dbqueries

import (
	"database/sql"
	"models"
	"github.com/davecgh/go-spew/spew"	
)

func (d DBQuery) CheckBalance(db *sql.DB, trans models.Transaction) (balance int) {

	var user models.User

	row := db.QueryRow("select balance from users where id = $1", trans.ID)

	err := row.Scan(&user.Balance)

	if err != nil {
		panic(err)
	}

	return user.Balance
}

func (d DBQuery) PostTrade (db *sql.DB, trans models.Transaction) error {

	sqlComm := "insert into transactions (userID, TYPE, SYMBOL, QUANTITY, PRICE) values ($1, $2, $3, $4, $5) returning id;"

	err := db.QueryRow(sqlComm, trans.ID, trans.Type, trans.Symbol, trans.Quantity, trans.Price).Scan(&trans.ID)

	return err
}

func (d DBQuery) UpdateBalance (db *sql.DB, id int, balance int) error {

	var user models.User

	sqlComm := "update users set balance = $1 where ID = $2 returning balance;"

	row := db.QueryRow(sqlComm, balance, id)

	err := row.Scan(&user.Balance)

	spew.Dump(user.Balance)

	return err
}

func (d DBQuery) VerifyHolding (db *sql.DB, trans Transaction) bool {

	var trans models.Transaction

	buyOrders := "select ID from transactions where userID = $1 and Symbol = '$2' and TYPE = 'Buy';"
	sellOrders := "select ID from transactions where userID = $1 and Symbol = '$2' and TYPE = 'Sell';"

	buyRows, err := db.Query(buyOrders, trans.userID, trans.Symbol, trans.Type)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// err := row.Scan(&user.Balance)

	spew.Dump(user.Balance)

	return err
}

