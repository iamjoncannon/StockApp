package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"models"
	"net/http"
	"utils"

	"dbqueries"

	"golang.org/x/crypto/bcrypt"
)

func (c Controller) LogIn(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)

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

			} else {
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

		returnData := make(map[string]string)

		stringBalance := fmt.Sprintf("%f", user.Balance)

		returnData["token"] = token
		returnData["Name"] = user.Name
		returnData["Balance"] = stringBalance

		j, err := json.Marshal(returnData)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, string(j))
	}
}
