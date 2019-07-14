package main

import (
	"log"
	"net/http"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	
	"driver"
	"controllers"
)

var db *sql.DB

func main(){

	gotenv.Load()

	controller := controllers.Controller{}

	db = driver.SQLConnect()
	defer db.Close()

	router := mux.NewRouter()

	go router.HandleFunc("/signup", controller.SignUp(db)).Methods("POST")
	go router.HandleFunc("/login", controller.LogIn(db)).Methods("POST")
	go router.HandleFunc("/maketransaction", controllers.TokenVerifyMiddleWare(controller.ConductTransaction(db))).Methods("POST")
	go router.HandleFunc("/getportfolio", controllers.TokenVerifyMiddleWare(controller.GetPortfolio(db))).Methods("POST")
	go router.HandleFunc("/getallTransactions", controllers.TokenVerifyMiddleWare(controller.GetTransactionHistory(db))).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Println("Listening on part 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
