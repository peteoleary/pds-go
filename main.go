package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Model[T any] interface {
	database_location() string
	table_name() string
	load_data(row *sql.Row) T
}

func open_database[T any](model Model[T]) *sql.DB {
	db, err := sql.Open("sqlite3", model.database_location())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func find_by_key[T any](model Model[T], db *sql.DB, key string, value string) T {
	sqlStmt := fmt.Sprintf("select * from %s where %s = ?", model.table_name(), key)
	row := db.QueryRow(sqlStmt, value)

	return model.load_data(row)
}

type Account struct {
	did              string
	email            string
	passwordScrypt   string
	emailConfirmedAt sql.NullTime
	invitesDisabled  int
}

func (a Account) database_location() string {
	return "../pds-data/account.sqlite"
}

func (a Account) table_name() string {
	return "account"
}
func (a Account) load_data(row *sql.Row) Account {
	err := row.Scan(&a.did, &a.email, &a.passwordScrypt, &a.emailConfirmedAt, &a.invitesDisabled)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

type Actor struct {
	did           string
	handle        string
	createdAt     string
	takedownRef   sql.NullString
	deactivatedAt sql.NullTime
	deleteAfter   sql.NullTime
}

func (a Actor) database_location() string {
	return "../pds-data/account.sqlite"
}

func (a Actor) table_name() string {
	return "actor"
}
func (a Actor) load_data(row *sql.Row) Actor {
	err := row.Scan(&a.did, &a.handle, &a.createdAt, &a.takedownRef, &a.deactivatedAt, &a.deleteAfter)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func main() {

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
