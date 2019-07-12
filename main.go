// export GOPATH=/Users/jonathancannon/projects/goTute/dino
// export GOBIN=/Users/jonathancannon/projects/goTute/dino/bin

package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
)

// whats with this json formatting?
// https://gobyexample.com/json

type User struct {
	ID 			int 	`json:"id"`
	Email		string 	`json:"email`
	Password	string 	`json:"password"`
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
  user     = "jonathancannon"
  dbname   = "test"
)

func main(){

	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "dbname=%s sslmode=disable",
    host, port, user, dbname)

	// pgUrl, err := pq.ParseURL("http://localhost:5432")

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

func signup(w http.ResponseWriter, r *http.Request) {

	r.Body()
	
	fmt.Println("signup invoked ", r)
	w.Write([]byte("successfully called signup\n"))
}

func login(w http.ResponseWriter, r *http.Request) {

	fmt.Println("login invoked")
}

func protectedEndpoint(w http.ResponseWriter, r *http.Request) {

}

func TokenVerifyMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return nil
}

