package connection

import (
	"database/sql"
	"fmt"
	"log"

	// test
	_ "github.com/lib/pq"
)

// Connect function receives database credentials and returns a connection to the database.
func Connect(dbEngine, user, password, dbName, host, port, sslmode string) (*sql.DB, error) {
	// Databse Connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", user, password, dbName, host, port, sslmode)
	db, err := sql.Open(dbEngine, connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Connection established")

	return db, nil
}
