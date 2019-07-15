package controllers

import (
	"database/sql"
	"net/http"
	"fmt"
	"encoding/json"

	// "github.com/davecgh/go-spew/spew"	

	"dbqueries"
	"models"
	"utils"
)

func (c Controller) GetTransactionHistory (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		// using the id in the actual header to make this 
		// request, otherwise a hacker could put someone
		// else's id in a post request and get their holdings

		decryptedid := r.Header["Decryptedid"][0]

		query := dbqueries.DBQuery{}

		rows, err := query.GetAllTransactions(db, decryptedid)

		var trans models.Transaction
		var transID int 
		var dateConducted string
		var errorObj models.Error
		
		type TransactionItem struct {
			Symbol		string 
			Quantity	float32 	
			Date 		string
		}

		transactionHistory := make(map[string]TransactionItem)

		defer rows.Close()
		
		for rows.Next() {

						// unique ID
						// of item in
						// Holding 	  // user's ID
			err = rows.Scan(&transID, &trans.ID, &trans.Type, &trans.Symbol, &trans.Quantity, &trans.Price, &dateConducted)

			if err != nil {

				fmt.Println(err)
				errorObj.Message = err.Error()
				utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
				return
			}

			// spew.Dump(transID, trans.ID, trans.Symbol, trans.Quantity, dateConducted)

			transactionHistory[fmt.Sprint(transID)] = TransactionItem{ Symbol: trans.Symbol, Quantity: trans.Quantity, Date: dateConducted }
		}

		j, err := json.Marshal(transactionHistory)

		utils.ResponseJSON(w, string(j))
	}
}	
