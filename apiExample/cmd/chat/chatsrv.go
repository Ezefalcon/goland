package main

import (
	"flag"
	"fmt"
	"github.com/ezefalcon/goland/internal/config"
	"github.com/ezefalcon/goland/internal/database"
	chat "github.com/ezefalcon/goland/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

func main() {
	cfg := initConfig()

	db, err := database.NewDatabase(cfg)
	defer db.Close()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//fmt.Println(cfg.DB.Driver)
	//fmt.Println(cfg.Version)

	service, _ := chat.New(db, cfg)
	httpService := chat.NewHTTPTransport(service)

	r := gin.Default()

	httpService.Register(r)
	r.Run()

	all := service.FindAll()
	for _, m := range all {
		fmt.Println(m)
	}
}

func createSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS messages (
		id integer primary key autoincrement,
		text varchar);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	// or, you can use MustExec, which panics on error
	insertMessage := `INSERT INTO messages (text) VALUES (?)`
	s := fmt.Sprintf("Message number %v", time.Now().Nanosecond())
	db.MustExec(insertMessage, s)
	return nil
}

func initConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
