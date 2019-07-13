package controllers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"database/sql"

	"models"
	"utils"
	"dbqueries"
	
	"github.com/davecgh/go-spew/spew"	
	"golang.org/x/crypto/bcrypt"
)


// we need access to the database, but that's defined
// in the driver file- solution - export the controller
// package as a struct, then define as the
// receiver of the SignUp handler function
// instantiate the struct in the main package to expose
// the handler function

type Controller struct{}

func (c Controller) SignUp (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		var user models.User 
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)
		
		fmt.Println("signup")

		spew.Dump(user)

		if user.Name == "" {
			error.Message = "Name not found in request"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Email == "" {
			error.Message = "Email not found in request"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		if user.Password == "" {
			error.Message = "No password in request"
			utils.RespondWithError(w, http.StatusBadRequest, error)
			return
		}

		hash, err := bcrypt.GenerateFromPassword( []byte(user.Password), 10 )

		if err != nil {
			panic(err)
		}

		query := dbqueries.DBQuery{}
				
		user, err = query.SignUp(db, user, string(hash))

		if err != nil {
		
			spew.Dump("there was a server error: ", err)
			error.Message = "Server error."
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		utils.ResponseJSON(w, user)
	}
}



