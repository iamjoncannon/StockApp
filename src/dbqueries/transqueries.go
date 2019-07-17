// #
package dbqueries

import (
	"database/sql"
	"models"
)

// pass in the entire model for these calls- the requirements for the database call could alter
// better to reduce that interface site by eliminating additional arguments required to make the call

func (d DBQuery) CheckBalance(db *sql.DB, trans models.Transaction) (balance float32, err error) {

	var user models.User

	row := db.QueryRow("select balance from users where id = $1", trans.ID)

	err = row.Scan(&user.Balance)

	return user.Balance, err
}

func (d DBQuery) GetCurrentHolding(db *sql.DB, trans models.Transaction) (balance float32, err error) {

	var holding models.Holding

	selectCall := "select current_holding from holdings where userid = $1 and symbol = $2"

	err = db.QueryRow(selectCall, trans.ID, trans.Symbol).Scan(&holding.Quantity)

	return holding.Quantity, err
}

func (d DBQuery) MakeFirstPurchase(db *sql.DB, trans models.Transaction) (newQuant float32, err error) {

	insertCall := "insert into holdings (userID, SYMBOL, CURRENT_HOLDING) values ($1, $2, $3) returning CURRENT_HOLDING;"

	err = db.QueryRow(insertCall, trans.ID, trans.Symbol, trans.Quantity).Scan(&trans.Quantity)

	return trans.Quantity, err
}

func (d DBQuery) UpdateHolding(db *sql.DB, newSum float32, trans models.Transaction) (newQuant float32, err error) {

	var holding models.Holding

	updateCall := "update holdings set current_holding = $1 where userid = $2 and symbol = $3 returning current_holding;"

	err = db.QueryRow(updateCall, newSum, trans.ID, trans.Symbol).Scan(&holding.Quantity)

	return holding.Quantity, err
}

func (d DBQuery) PostTrade(db *sql.DB, trans models.Transaction) error {

	sqlComm := "insert into transactions (userID, TYPE, SYMBOL, QUANTITY, PRICE) values ($1, $2, $3, $4, $5) returning id;"

	err := db.QueryRow(sqlComm, trans.ID, trans.Type, trans.Symbol, trans.Quantity, trans.Price).Scan(&trans.ID)

	return err
}

func (d DBQuery) UpdateBalance(db *sql.DB, id int, balance float32) (updatedBalance float32, err error) {

	var user models.User

	sqlComm := "update users set balance = $1 where ID = $2 returning balance;"

	err = db.QueryRow(sqlComm, balance, id).Scan(&user.Balance)

	return user.Balance, err
}

/// this is the prepared statement version of the transaction:

/*

func (d DBQuery) FullTradeTransaction(db *sql.DB, newSum float32, userId int, balance float32, trans models.Transaction) (updatedBalance float32, newQuant float32, err error) {

	var holding models.Holding
	var user models.User

	updateHoldingsCall := "update holdings set current_holding = $1 where userid = $2 and symbol = $3 returning current_holding;"

	err = db.QueryRow(updateHoldingsCall, newSum, trans.ID, trans.Symbol).Scan(&holding.Quantity)

	insertTransactionCall := "insert into transactions (userID, TYPE, SYMBOL, QUANTITY, PRICE) values ($1, $2, $3, $4, $5) returning id;"

	err := db.QueryRow(insertTransactionCall, trans.ID, trans.Type, trans.Symbol, trans.Quantity, trans.Price).Scan(&trans.ID)

	updateBalanceCall := "update users set balance = $1 where ID = $2 returning balance;"

	err = db.QueryRow(updateBalanceCall, balance, userId).Scan(&user.Balance)

	return holding.Quantity, user.Balance, err
}

*/

func (d DBQuery) FullTradeTransaction(db *sql.DB, newHoldingAmount float32, userId int, newBalance float32, trans models.Transaction) error {

	tx, err := db.Begin()

	if err != nil {

		return err
	}

	{
		updateHoldingsCall := "update holdings set current_holding = $1 where userid = $2 and symbol = $3 returning current_holding;"

		stmt, err := tx.Prepare(updateHoldingsCall)

		if err != nil {
			err := tx.Rollback()
			return err
		}

		defer stmt.Close()

		_, err = stmt.Exec(newHoldingAmount, trans.ID, trans.Symbol)

		if err != nil {
			err := tx.Rollback()
			return err
		}
	}

	{
		insertTransactionCall := "insert into transactions (userID, TYPE, SYMBOL, QUANTITY, PRICE) values ($1, $2, $3, $4, $5) returning id;"

		stmt, err := tx.Prepare(insertTransactionCall)

		if err != nil {
			err := tx.Rollback()
			return err
		}

		defer stmt.Close()

		if _, err := stmt.Exec(trans.ID, trans.Type, trans.Symbol, trans.Quantity, trans.Price); err != nil {
			err := tx.Rollback()
			return err
		}
	}

	{
		updateBalanceCall := "update users set balance = $1 where ID = $2 returning balance;"

		stmt, err := tx.Prepare(updateBalanceCall)

		if err != nil {
			err := tx.Rollback()
			return err
		}

		defer stmt.Close()

		if _, err := stmt.Exec(newBalance, userId); err != nil {
			err := tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	return err
}
