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
	var account_db = open_database(account)
	if account_db == nil {
		log.Fatal("Failed to open database")
	}
	defer account_db.Close()

	accountData := find_by_key(account, account_db, "email", "bksy@timelight.com")
	fmt.Printf("Account: %+v\n", accountData)

	actorData := find_by_key(Actor{}, account_db, "handle", "pete.bsky.timelight.com")
	fmt.Printf("Actor: %+v\n", actorData)

	record := Record{}
	record.actor_did = "did:plc:2yn32k65auyhjo2thnya3hlg"
	record_db := open_database(record)

	if record_db == nil {
		log.Fatal("Failed to open store database")
	}
	defer record_db.Close()
	recordData := find_by_key(record, record_db, "uri", "at://did:plc:2yn32k65auyhjo2thnya3hlg/app.bsky.feed.post/3lhab2cuszs22")
	fmt.Printf("Record: %+v\n", recordData)
}
