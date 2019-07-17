package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"dbqueries"
	"models"
	"utils"

	"golang.org/x/crypto/bcrypt"
)

type Controller struct{}

func (c Controller) SignUp(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

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

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

		if err != nil {
			panic(err)
		}

		query := dbqueries.DBQuery{}

		user, err = query.SignUp(db, user, string(hash))

		if err != nil {

			fmt.Println("there was a server error: ", err)
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		token, err := GenerateToken(user)

		if err != nil {

			fmt.Println("there was a server error: ", err)
			error.Message = err.Error()
			utils.RespondWithError(w, http.StatusInternalServerError, error)
			return
		}

		w.WriteHeader(http.StatusOK)

		utils.ResponseJSON(w, token)
	}
}
