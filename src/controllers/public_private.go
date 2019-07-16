// #
package controllers

import (
	"os"
	"log"
	"net/http"
	"strings"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	
	"models"
	"utils"
)

func GenerateToken(user models.User)(string, error){

	secret := os.Getenv("PRIVATE_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"id": user.ID,
		"iss": "broker",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {

													// this refers to the original http
													// request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

		var errorObj models.Error

		// process the token from the header
		authHeader := r.Header.Get("Authorization")
		tokenToVerify := strings.TrimSpace(string(strings.Split(authHeader, "Bearer")[1]))

		// parsing the token we received- method expects unvalidated token 
		// and a callback used to verify the token, which returns the private key
		verifiedToken, error := jwt.Parse(tokenToVerify, func(token *jwt.Token)(interface{}, error) {
			
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error with your token")
			}

			// return []byte("this is my private key"), nil
			return []byte(os.Getenv("PRIVATE_KEY")), nil
		})

		if error != nil {

			errorObj.Message = error.Error()
			utils.RespondWithError(w, http.StatusUnauthorized, errorObj)
		}

		if verifiedToken.Valid {

			// we're going to "decrypt" the email and put it into
			// the request header for the next function call
			// these are called type assertions- 
			// https://tour.golang.org/methods/15
			claims, _ := verifiedToken.Claims.(jwt.MapClaims)
		
			floatId, _ := claims["id"].(float64)

			stringId := fmt.Sprintf("%d", int(floatId) )

			r.Header.Set("decryptedId",  stringId)

			// this just curries into 
			// whatever the argument function is
			next.ServeHTTP(w, r)

		} else { 
		
			errorObj.Message = error.Error()
			utils.RespondWithError(w, http.StatusUnauthorized, errorObj)
			return
		}
	})
}
