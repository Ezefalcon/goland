package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sqlx.DB
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	createSchema(db)
}

// User
type User struct {
	ID int `db:"id"`
	Name string `db:"name"`
}

func createSchema(db *sqlx.DB) {
	schema := `CREATE TABLE user (id integer primary key AUTOINCREMENT,
				 name varchar(56));`
	_, err := db.Exec(schema)
	if err != nil {
		panic(err.Error())
	}
	db.MustExec("INSERT INTO user (name) VALUES (?)", "Jane Doe")
	rows, err := db.Queryx("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var u User
		err = rows.StructScan(&u)
		fmt.Println(u)
	}
}
