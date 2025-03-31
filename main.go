package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	account := Account{}
	var db = open_database(account)
	if db == nil {
		log.Fatal("Failed to open database")
	}
	defer db.Close()

	accountData := find_by_key(account, db, "email", "bksy@timelight.com")
	fmt.Printf("Account: %+v\n", accountData)

	actorData := find_by_key(Actor{}, db, "handle", "pete.bsky.timelight.com")
	fmt.Printf("Actor: %+v\n", actorData)
}
