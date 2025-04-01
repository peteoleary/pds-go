package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Model[T any] interface {
	database_location() string
	table_name() string
	load_data(row *sql.Rows) T
}

func open_database[T any](model Model[T]) *sql.DB {
	var location = model.database_location()
	log.Printf("Opening database at %s\n", location)
	db, err := sql.Open("sqlite3", location)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func find_by_key[T any](model Model[T], db *sql.DB, key string, value string) T {
	sqlStmt := fmt.Sprintf("select * from %s where %s = ?", model.table_name(), key)
	row, err := db.Query(sqlStmt, value)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	if !row.Next() {
		log.Fatal("No rows returned")
	}

	return model.load_data(row)
}

func find_all[T any](model Model[T], db *sql.DB) []T {
	sqlStmt := fmt.Sprintf("select * from %s", model.table_name())
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var results []T
	for rows.Next() {
		row := model.load_data(rows)
		results = append(results, row)
	}
	return results
}

type Account struct {
	did              string
	email            string
	passwordScrypt   string
	emailConfirmedAt sql.NullTime
	invitesDisabled  int
}

func (a Account) database_location() string {
	return fmt.Sprintf("%s/account.sqlite", os.Getenv("PDS_DATA_DIRECTORY"))
}

func (a Account) table_name() string {
	return "account"
}
func (a Account) load_data(row *sql.Rows) Account {
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
	return fmt.Sprintf("%s/account.sqlite", os.Getenv("PDS_DATA_DIRECTORY"))
}

func (a Actor) table_name() string {
	return "actor"
}
func (a Actor) load_data(row *sql.Rows) Actor {
	err := row.Scan(&a.did, &a.handle, &a.createdAt, &a.takedownRef, &a.deactivatedAt, &a.deleteAfter)
	if err != nil {
		log.Fatal(err)
	}
	return a
}

func hash_did(did string) string {
	data := []byte(did)
	hash := sha256.New()
	hash.Write(data)
	hashInBytes := hash.Sum(nil)
	hashString := fmt.Sprintf("%x", hashInBytes)
	return hashString
}

func (a Actor) get_actor_directory() string {
	hash := hash_did(a.did)
	directory := fmt.Sprintf("%s/%s/%s", os.Getenv("PDS_ACTOR_STORE_DIRECTORY"), hash[:2], a.did)
	return directory
}

type Record struct {
	actor_did   string // set this to the actor's did
	uri         string
	cid         string
	collection  string
	rkey        string
	repoRev     string
	indexedAt   string
	takedownRef sql.NullString
}

func (r Record) database_location() string {
	return fmt.Sprintf("%s/store.sqlite", Actor{did: r.actor_did}.get_actor_directory())
}
func (r Record) table_name() string {
	return "record"
}
func (r Record) load_data(row *sql.Rows) Record {
	err := row.Scan(&r.uri, &r.cid, &r.collection, &r.rkey, &r.repoRev, &r.indexedAt, &r.takedownRef)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
