package controllers

import (
	"database/sql"
	"net/http"
	"fmt"
	"encoding/json"

	"dbqueries"
	"models"
	"utils"
)

func (c Controller) GetPortfolio (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		// using the id in the actual header to make this 
		// request, otherwise a hacker could put someone
		// else's id in a post request and get their holdings

		decryptedid := r.Header["Decryptedid"][0]

		query := dbqueries.DBQuery{}

		rows, err := query.GetPortfolio(db, decryptedid)

		var holding models.Holding
		var holdingID int 
		var errorObj models.Error

		portfolio := make(map[string]models.Holding)

		defer rows.Close()
		
		for rows.Next() {

						// unique ID
						// of item in
						// Holding 	  // user's ID
			err = rows.Scan(&holdingID, &holding.ID, &holding.Symbol, &holding.Quantity)

			if err != nil {

				fmt.Println(err)
				errorObj.Message = err.Error()
				utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
				return
			}

			portfolio[fmt.Sprint(holding.Symbol)] = models.Holding{ ID: holding.ID, Symbol: holding.Symbol, Quantity: holding.Quantity }
		}

		j, err := json.Marshal(portfolio)

		utils.ResponseJSON(w, string(j))
	}
}	
