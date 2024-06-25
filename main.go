package main

import (
	"database/sql"
	// "fmt"
	"log"

	db "github.com/ekefan/backend-skudoosh/internal/db/sqlc"
	api "github.com/ekefan/backend-skudoosh/internal/server"
	"github.com/ekefan/backend-skudoosh/internal/utils"
	_ "github.com/lib/pq"
)

func main() {
	// load environnent variables
	config, err :=  utils.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(dbConn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	// fmt.Println(config.ServerAddress)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}

}
