package database

import (
	"database/sql"
	"log"

	_ "gopkg.in/cq.v1"
)

var db *sql.DB

/**
 * Setup opens a connection to the database and pings
 * for user
 */
func Setup() {
	db, _ = sql.Open("neo4j-cypher", "http://neo4j:testings@localhost:7474")
	// defer db.Close()

	err := db.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Print("Database ready for queries.")
}

/**
 * Get returns DB instance
 */
func Get() *sql.DB {
	return db
}
