package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"cache"
	"controllers"
	"driver"
)

var db *sql.DB

func main() {

	gotenv.Load()

	db = driver.SQLConnect()
	defer db.Close()

	pool := cache.NewPool()
	conn := pool.Get()
	defer conn.Close()
	err := cache.Ping(conn)

	if err != nil {
		fmt.Println(err)
	}

	controller := controllers.Controller{}
	router := mux.NewRouter()

	go router.HandleFunc("/signup", controller.SignUp(db)).Methods("POST")
	go router.HandleFunc("/login", controller.LogIn(db)).Methods("POST")
	go router.HandleFunc("/getportfolio", controllers.TokenVerifyMiddleWare(controller.GetPortfolio(db))).Methods("POST")
	go router.HandleFunc("/getallTransactions", controllers.TokenVerifyMiddleWare(controller.GetTransactionHistory(db))).Methods("POST")
	go router.HandleFunc("/maketransaction", controllers.TokenVerifyMiddleWare(controller.ConductTransaction(db))).Methods("POST")
	go router.HandleFunc("/ohlc/{symbol}", controllers.TokenVerifyMiddleWare(controllers.FetchOpeningPrice(conn))).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	log.Println("Listening on part 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}
