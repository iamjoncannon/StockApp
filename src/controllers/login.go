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
	"github.com/davecgh/go-spew/spew"	
	"golang.org/x/crypto/bcrypt"
)

func (c Controller) LogIn (db *sql.DB) http.HandlerFunc {

	return func (w http.ResponseWriter, r *http.Request) {
		
		var user models.User 
		var error models.Error

		json.NewDecoder(r.Body).Decode(&user)
		
		// spew.Dump(user)

		// fmt.Println("this is the login callback: ", user)

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

		returnData := make(map[string]string)

		// type ReturnData struct {

		// 	returnToken string 	`json:"returnToken"` 
		// 	Name 		string 	`json:"Name"`
		// 	Balance 	float32 `json:"Balance"`
		// }

		// I was having a hard time converting a struct to
		// a json- converting the map was working for me though
		// have to type coerce this and then coerce back on the client
		stringBalance := fmt.Sprintf("%f", user.Balance)

		spew.Dump(stringBalance)

		returnData["token"] = token
		returnData["Name"] = user.Name
		returnData["Balance"] = stringBalance

		// returnData := ReturnData{ returnToken: token, Name: user.Name, Balance: user.Balance}

		j, err := json.Marshal(returnData)

		if err != nil {
			fmt.Println(err)
		}

		spew.Dump(string(j))

		utils.ResponseJSON(w, string(j))
	}
}