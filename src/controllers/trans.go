package controllers

import (
	"net/http"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	// "github.com/davecgh/go-spew/spew"	

	"models"
	"dbqueries"
	"utils"
)

func (c Controller) ConductTransaction (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){
	
		var trans models.Transaction 
		var newBalance float32
		var newHoldingAmount float32
		var errorObj models.Error

		json.NewDecoder(r.Body).Decode(&trans)

		// populate the ID for the transaction with the 
		// id in the client's token
		parsedID, _ := strconv.ParseInt(r.Header["Decryptedid"][0], 10, 64)
		trans.ID = int(parsedID) 

		query := dbqueries.DBQuery{}

		// in order for the transaction to succeeed, each part needs
		// to clear, if not, we need to fail the whole transaction
		// this means any error needs to dead the whole function call,
		// and each database call must return an error object

		// to handle errors we "respond with error" with the correct
		// http status 

		// spew.Dump(trans)

		currentBalance, err := query.CheckBalance(db, trans)

		if err != nil {

			fmt.Println(err)
			errorObj.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
			return
		}

		fmt.Println("current balance: ", currentBalance)

		currentHolding, err := query.GetCurrentHolding(db, trans)

		if err != nil {

			// buy call where its their first purchase will return this
			if err.Error() == "sql: no rows in result set" {

				// this is a separate database transaction- 

				firstPurchase, err := query.MakeFirstPurchase(db, trans)

				if err != nil {

					fmt.Println(err)
					errorObj.Message = err.Error()
					utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
					return
				}

				if firstPurchase != trans.Quantity {

					fmt.Println("error conducting first stock purchase")
					errorObj.Message = "internal server error"
					utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
					return
				}
			} else {

				errorObj.Message = err.Error()
				utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
				return
			}
		}

		fmt.Println("current holding: ", currentHolding)

		if trans.Type == "Buy"{

			hasFundsToCoverTransaction := currentBalance > trans.Quantity * trans.Price 

			if !hasFundsToCoverTransaction {

				errorObj.Message = "Insufficient balance to cover trade"
				utils.ResponseJSON(w, errorObj)
				return
			}

			newBalance = currentBalance - ( trans.Quantity * trans.Price)
			newHoldingAmount = currentHolding + trans.Quantity

		}else{

			hasStockToCoverSale := currentHolding >= trans.Quantity

			if !hasStockToCoverSale {

				errorObj.Message = "You do not hold stock to cover that sell order"
				utils.ResponseJSON(w, errorObj)
				return
			}

			newBalance = currentBalance + trans.Quantity * trans.Price
			newHoldingAmount = currentHolding - trans.Quantity
		}

		fmt.Println("newBalance: ", newBalance)
		fmt.Println("new Holding: ", newHoldingAmount)

		// this next part raises an interesting strategic question:
		// which party should be given more protections by the
		// transaction process?

		// if you post the trade before the balance is updated, then
		// the person may make multiple buy orders that could clear
		// before the database is updated to reflect that they have
		// a negative balance- is that a big deal? You could just
		// hit their credit card up again later...

		// but, if you post a sell order, and the person doesn't have
		// the holdings to cover it, then you're on the hook for the
		// counter party, you have to cover the stock that you said you
		// would sell them- point being, always update the holding before
		// posting the trade

		// otherwise could get an off-by-one type error- if they post one more
		// sale than they have, on the next to last trade, the trade
		// would clear, then the gap between that trade clearing and
		// the holdings updating, the person could make another sale
		// even though they couldn't cover the trade

		returnedHolding, err := query.UpdateHolding(db, newHoldingAmount, trans)

		if err != nil {

			fmt.Println(err)
			errorObj.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
			return
		}

		if returnedHolding != newHoldingAmount {

			fmt.Println("error updating holding")
			errorObj.Message = "internal server error"
			utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
			return
		}

		// if the holding update is the same as what you thought
		// it would be above

		err = query.PostTrade(db, trans)

		if err != nil {

			fmt.Println(err)
			errorObj.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
			return
		}

		returnedBalance, err := query.UpdateBalance(db, trans.ID, newBalance)

		if returnedBalance != newBalance {

			fmt.Println("error updating balance")
			errorObj.Message = "internal server error"
			utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
			return
		}
	}
}

