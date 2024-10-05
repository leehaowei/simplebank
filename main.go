package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/leehaowei/simplebank/api"
	db "github.com/leehaowei/simplebank/db/sqlc"
	_ "github.com/lib/pq"
)

const serverAddress = "0.0.0.0:8080"

var (
	dbDriver = FetchEnv("GOOSE_DRIVER")
	dbSource = FetchEnv("DATABASE_URL")
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func FetchEnv(key string) string {
	value, exists := os.LookupEnv(key)

	if !exists {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}

	return value
}
