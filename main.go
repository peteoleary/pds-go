package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Model[T any] interface {
	table_name() string
	load_data(row *sql.Row) T
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
	createdAt     sql.NullTime
	takedownRef   string
	deactivatedAt sql.NullTime
	deleteAfter   sql.NullTime
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

	db, err := sql.Open("sqlite3", "../pds-data/account.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	accountData := find_by_key(Account{}, db, "email", "bksy@timelight.com")
	fmt.Printf("Account: %+v\n", accountData)
}
