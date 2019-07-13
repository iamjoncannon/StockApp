package dbqueries

import (
	"database/sql"
	"models"
	"fmt"
	"github.com/davecgh/go-spew/spew"	
)

func stopThisErrorHere(){
	fmt.Println("The Google engineers really want these programs to be as efficient as possible.")
	spew.Dump("To that end, type safety is tedious but important")
}

type DBQuery struct {}

func (d DBQuery) SignUp(db *sql.DB, user models.User, password string) (models.User, error) {

	sqlComm := "insert into users (name, email, password) values ($1, $2, $3) RETURNING id;"

	err := db.QueryRow(sqlComm, user.Name, user.Email, password).Scan(&user.ID)
	
	user.Password = "<ENCRYPTED>"

	return user, err 
}

func (d DBQuery) LogIn(db *sql.DB, user models.User)(models.User, error) {

	row := db.QueryRow("select * from users where email = $1", user.Email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Balance)

	return user, err
}

func (d DBQuery) GetPortfolio(db *sql.DB, id string ) (rows *sql.Rows, err error) {

	rows, err = db.Query("select * from holdings where userid = $1", id)

	return rows, err
}

func (d DBQuery) GetAllTransactions(db *sql.DB, id string ) (rows *sql.Rows, err error) {

	rows, err = db.Query("select * from transactions where userid = $1", id)

	return rows, err
}

