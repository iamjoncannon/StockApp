package dbqueries

import (
	"database/sql"
	"models"
)

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

/*

"Prepared statement"- either all three database calls succeed or all three will be failed.

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

	// only if all three clear will the
	// transaction be committed

	err = tx.Commit()

	return err
}
