package controllers

import (
	// "os"
	// "log"
	"net/http"
	"database/sql"
	"encoding/json"
	// "strings"
	// "fmt"

	// jwt "github.com/dgrijalva/jwt-go"
	// "github.com/davecgh/go-spew/spew"	
	
	"models"
	"dbqueries"
	"utils"
)

func (c Controller) ConductTransaction (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request){
	
		var trans models.Transaction 
		var error models.Error
		var newBalance int

		json.NewDecoder(r.Body).Decode(&trans)

		// spew.Dump(trans)

		query := dbqueries.DBQuery{}

		balance := query.CheckBalance(db, trans)

		if trans.Type == "buy"{

			hasFundsToCoverTransaction := balance > trans.Quantity * trans.Price 

			if !hasFundsToCoverTransaction{

				error.Message = "Insufficient balance to cover trade"
				utils.ResponseJSON(w, error)
				return
			}

			newBalance = balance - (trans.Quantity * trans.Price)

		}else{

			// holdsStocksToCoverSale := 

			newBalance = balance + trans.Quantity * trans.Price

		}

		err := query.PostTrade(db, trans)

		if err != nil {
			panic(err)
		}

		// spew.Dump(balance)

		err = query.UpdateBalance(db, trans.ID, newBalance)

		if err != nil {
			panic(err)
		}

	}
}

