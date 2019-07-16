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


	// why define the callback inside TokenVerify as a method on a controller struct?
	// otherwise you will lose reference to the db pointer
	// When the Token middleware gets called, it is passed r and w by handlefunc.
	// The argument signature of mux' Handlefunc method is "http.HandlerFunc", meaning
	// we can't add a db pointer variable as a argument to the middleware, otherwise
	// the compiler would reject it for not having valid typing

	// when TokenVal.. resolves it calls its argument function with r and w, but the argument function
	// now exists in the controllers package namespace. Because its call site is no longer
	// the main package, therefore, if the method were invoked directly, as the controllers above, 
	// it would in essence not really be "imported" from that package namespace at all. 
	// Defining Token's callback as a struct inside this package means it has reference
	// to this package's variable environment, including the db pointer

	// this isn't a problem with the unprotected routes, because those are being invoked
	// inside this namespace

	// another technique would be to somehow export the db pointer to the other namespace
	// and use it there within the callbacks

	go router.HandleFunc("/getportfolio", controllers.TokenVerifyMiddleWare(controller.GetPortfolio(db))).Methods("POST")
	go router.HandleFunc("/getallTransactions", controllers.TokenVerifyMiddleWare(controller.GetTransactionHistory(db))).Methods("POST")
	go router.HandleFunc("/maketransaction", controllers.TokenVerifyMiddleWare(controller.ConductTransaction(db))).Methods("POST")
	go router.HandleFunc("/ohlc/{symbol}", controllers.TokenVerifyMiddleWare(controllers.FetchOpeningPrice())).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Println("Listening on part 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}