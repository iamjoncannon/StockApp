package controllers

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"
	"models"
	"utils"
	"fmt"

	"dbqueries"
	
	"golang.org/x/crypto/bcrypt"
)

func (c Controller) LogIn (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {

		var user models.User 
		var jwt models.JWT
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

		fmt.Println(user)

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

		// store submitted password to check against decrypted
		// password in the database 

		password := user.Password 

		query := dbqueries.DBQuery{}
		
		user, err := query.LogIn(db, user)

		if err != nil {

			if err == sql.ErrNoRows {
				
				error.Message = "User credentials not found in system"
				utils.RespondWithError(w, http.StatusBadRequest, error)
				return
			}else{
				log.Fatal(err)
			}
		}

		hashedPassword := user.Password 

		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

		if err != nil {

			error.Message = "Invalid Password"
			utils.RespondWithError(w, http.StatusUnauthorized, error)
			return 
		}

		token, err := GenerateToken(user)

		if err != nil {
			log.Fatal(err)
		}

		w.WriteHeader(http.StatusOK)

		jwt.Token = token

		utils.ResponseJSON(w, jwt)
	}
}