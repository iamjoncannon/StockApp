package controllers

import (
	"os"
	"log"
	"net/http"
	"strings"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	// "github.com/davecgh/go-spew/spew"	
	
	"models"
	"utils"
)

func GenerateToken(user models.User)(string, error){

	secret := os.Getenv("PRIVATE_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
		"iss": "broker",
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}


// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImN1dHR5QGV4YW1wbGUuY29tIiwiaXNzIjoiYnJva2VyIn0.sFCN6BlhfHFkGwSWi1aTPpLOBTU9UnS_dTQ1pRPeg2I" http://localhost:3000/protected

// this is an example of currying- the verification function receives the 
// handler and, through the mux package, the entire original function is
// itself actually invoked at the end of the call
// like, the key is to add () after the mux.router calls- everything
// after return below is being called, and the w and r arguments
// refer to the original http request with the token that's being
// verified- this isn't honestly 100% clear below - 

// whats happening is.. router.HandleFunc takes an endpoint and a callback,
// when the endpoint is hit it mounts the call back to the route and passes in
// the original http request object 

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

			return []byte("this is my private key"), nil
		})

		// spew.Dump(verifiedToken, error)

		if error != nil {

			errorObj.Message = error.Error()
			utils.RespondWithError(w, http.StatusUnauthorized, errorObj)
		}

		if verifiedToken.Valid {

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
