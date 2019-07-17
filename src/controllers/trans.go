// #

package controllers

import (
	"database/sql"
	"dbqueries"
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"strconv"
	"utils"
)

func (c Controller) ConductTransaction(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

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

		currentBalance, err := query.CheckBalance(db, trans)

		if err != nil {

			fmt.Println(err)
			errorObj.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, errorObj)
			return
		}

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

		if trans.Type == "Buy" {

			hasFundsToCoverTransaction := currentBalance > trans.Quantity*trans.Price

			if !hasFundsToCoverTransaction {

				errorObj.Message = "Insufficient balance to cover trade"
				utils.ResponseJSON(w, errorObj)
				return
			}

			newBalance = currentBalance - (trans.Quantity * trans.Price)
			newHoldingAmount = currentHolding + trans.Quantity

		} else {

			hasStockToCoverSale := currentHolding >= trans.Quantity

			if !hasStockToCoverSale {

				errorObj.Message = "You do not hold stock to cover that sell order"
				utils.ResponseJSON(w, errorObj)
				return
			}

			newBalance = currentBalance + trans.Quantity*trans.Price
			newHoldingAmount = currentHolding - trans.Quantity
		}

		// we use a prepared statement to prevent errors in recording
		// this set of transactions- either they all suceeed or they
		// all fail

		err = query.FullTradeTransaction(db, newHoldingAmount, trans.ID, newBalance, trans)

		/*

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

		*/

		utils.ResponseJSON(w, newBalance)
	}
}
