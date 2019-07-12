package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"database/sql"
	"encoding/json"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/davecgh/go-spew/spew"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID 			int 	`json:"id"`
	Name 		string 	`json:"name"`
	Email		string 	`json:"email`
	Password	string 	`json:"password"`
	Balance		int 	`json:"balance"`
}

type JWT struct {
	Token string 	`json:"token"`
}

type Error struct {
	Message string `json:"message"`
}

// global variable available in the whole package
var db *sql.DB

// database connection info:
const (
  host     = "localhost"
  port     = 5432
  user     = "jcannon"
  dbname   = "test"
)

func main(){

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	fmt.Println("Established connection to database: ", dbname)

	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/signup", signup).Methods("POST")
	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/protected", TokenVerifyMiddleWare(protectedEndpoint)).Methods("GET")

	log.Println("Listening on part 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}

// test command:
// curl http://localhost:3000/signup -X POST -d '{"name":"Cutty","password":"yadayadayada", "email":"cutty@example.com"}'

func respondWithError(w http.ResponseWriter, status int, error Error ){
	// error response
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(error)
}

func responseJSON(w http.ResponseWriter, data interface{}) {

	json.NewEncoder(w).Encode(data)
}

func signup(w http.ResponseWriter, r *http.Request) {

	// instantiating structs defined above
	var user User 
	var error Error

	// https://stackoverflow.com/questions/38172661/what-is-the-meaning-of-and-in-golang
	// & returns the memory address of the following variable.

	// * returns the value of the following variable 
	// (which should hold the memory address of a variable, 
	// unless you want to get weird output and possibly problems 
	// because you're accessing your computer's RAM)

	// http://piotrzurek.net/2013/09/20/pointers-in-go.html

	// & in front of variable name is used to retrieve the address 
	// of where this variable’s value is stored. That address is what the pointer 
	// is going to store.

	// aka pass by reference- its a reference to that variable that exists elsewhere?
	// but ...aren't all variables references..... .... ....
	// maybe user is an object...


	json.NewDecoder(r.Body).Decode(&user)
	
	fmt.Println("signup callback invoked")
	fmt.Println("_______________________")

	spew.Dump(user)

	if user.Name == "" {
		error.Message = "Name not found in request"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Email == "" {
		error.Message = "Email not found in request"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "No password in request"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	hash, err := bcrypt.GenerateFromPassword( []byte(user.Password), 10 )

	if err != nil {
		panic(err)
	}

	fmt.Println("password: ", user.Password)
	fmt.Println("email:", user.Email, "name: ", user.Name, "password: ", string(hash))

	sqlComm := "insert into users (name, email, password) values ($1, $2, $3) RETURNING id;"

	err = db.QueryRow(sqlComm, user.Name, user.Email, string(hash)).Scan(&user.ID)
	
	if err != nil {
		spew.Dump("there was a server error: ", err)
		error.Message = "Server error."
		respondWithError(w, http.StatusInternalServerError, error)
		return
	}

	user.Password = ""
	w.Header().Set("Content-Type", "application/json")
	responseJSON(w, user)
	// w.Write([]byte("successfully called signup\n"))
}

func GenerateToken(user User)(string, error){

	var err error

	secret := "this is my private key"

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

// curl http://localhost:3000/login -X POST -d '{"name":"Cutty","password":"yadayadayada", "email":"cutty@example.com"}'

func login(w http.ResponseWriter, r *http.Request) {

	var user User 
	var jwt JWT
	var error Error

	json.NewDecoder(r.Body).Decode(&user)

	if user.Email == "" {
		error.Message = "Email not found in request"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	if user.Password == "" {
		error.Message = "No password in request"
		respondWithError(w, http.StatusBadRequest, error)
		return
	}

	// store submitted password to check against decrypted
	// password in the database 

	password := user.Password 

	row := db.QueryRow("select * from users where email = $1", user.Email)

	// Scan takes the row data and instantiates user struct

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Balance)

	if err != nil {

		if err == sql.ErrNoRows {
			error.Message = "User credentials not found in system"
			respondWithError(w, http.StatusBadRequest, error)
			return
		}else{
			log.Fatal(err)
		}
	}

	hashedPassword := user.Password 

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		error.Message = "Invalid Password"
		respondWithError(w, http.StatusUnauthorized, error)
		return 
	}

	token, err := GenerateToken(user)

	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)

	jwt.Token = token

	responseJSON(w, jwt)
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {

}

// curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImN1dHR5QGV4YW1wbGUuY29tIiwiaXNzIjoiYnJva2VyIn0.sFCN6BlhfHFkGwSWi1aTPpLOBTU9UnS_dTQ1pRPeg2I" http://localhost:3000/protected


func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		var errorObj Error

		// process the token from the header
		authHeader := r.Header.Get("Authorization")
		tokenToVerify := strings.TrimSpace(string(strings.Split(authHeader, "Bearer")[1]))

		// parsing the token we received- method expects unvalidated token 
		// and a callback used to verify the token, which returns the private key
		verifiedToken, error := jwt.Parse(tokenToVerify, func(token *jwt.Token)(interface{}, error) {
			
			// spew.Dump(token)

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error with your token")
			}

			return []byte("this is my private key"), nil
		})

		if error != nil {

			errorObj.Message = error.Error()
			respondWithError(w, http.StatusUnauthorized, errorObj)
		}

		if verifiedToken.Valid {

			// this just curries into 
			// whatever the argument function is
			next.ServeHttp(w, r)
		} else { 
			errorObj.Message = error.Error()
			respondWithError(w, http.StatusUnauthorized, errorObj)
			return
		}



	})
}

