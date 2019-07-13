package dbqueries

import (
	"database/sql"
	"models"
	// "github.com/davecgh/go-spew/spew"	
)

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


// func (d DBQuery) GetPortfolio(db *sql.DB, id int )[]int()

