package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"dbqueries"
	"models"
	"utils"
)

func (c Controller) GetTransactionHistory(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// using the id in the actual header to make this
		// request, otherwise someone could put someone
		// else's id in a post request and get their holdings

		decryptedid := r.Header["Decryptedid"][0]

		query := dbqueries.DBQuery{}

		rows, err := query.GetAllTransactions(db, decryptedid)

		var trans models.Transaction
		var transID int
		var dateConducted string
		var errorObj models.Error

		// have to define separately because we use the
		// database to timestamp the transaction- the Golang
		// and SQL models aren't exactly the same
		type TransactionItem struct {
			Symbol   string
			Quantity float32
			Date     string
			Price    float32
			Type     string
		}

		transactionHistory := make(map[string]TransactionItem)

		defer rows.Close()

		for rows.Next() {

			err = rows.Scan(&transID, &trans.ID, &trans.Type, &trans.Symbol, &trans.Quantity, &trans.Price, &dateConducted)

			if err != nil {

				fmt.Println(err)
				errorObj.Message = err.Error()
				utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
				return
			}

			transactionHistory[fmt.Sprint(transID)] = TransactionItem{Type: trans.Type,
				Symbol:   trans.Symbol,
				Quantity: trans.Quantity,
				Date:     dateConducted,
				Price:    trans.Price}
		}

		j, err := json.Marshal(transactionHistory)

		utils.ResponseJSON(w, string(j))
	}
}
