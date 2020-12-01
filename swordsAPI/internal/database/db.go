package database

import (
	"errors"
	"fmt"
	"github.com/Ezefalcon/golang/swordsAPI/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
)

// NewDatabase ...
func NewDatabase(conf *config.Config) (*sqlx.DB, error) {
	switch conf.DB.Type {
	case "sqlite3":
		db, err := sqlx.Open(conf.DB.Driver, conf.DB.Conn)
		if err != nil {
			return nil, err
		}

		db.Ping()
		if err != nil {
			return nil, err
		}
		return db, nil
	default:
		return nil, errors.New("invalid db type")
	}
}

func PopulateDatabase(db *sqlx.DB, filepath string) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}

	query := string(file)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CreateSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS sword (
		id integer primary key autoincrement,
		name varchar,
		speciality varchar,
		damage varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}