package Config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

/*
	Create database for the application
*/
func CreateDatabase() {
	file, err := os.Create("database.CryptoTracker")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}

/*
	Create bitcoin_tracker table to store bitcoin values
*/
func CreateBitcoinTrackerTable(db *sql.DB) {
	users_table := `CREATE TABLE bitcoin_tracker (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "CoinType" TEXT,
        "Price" INT,
        "Timestamp" DATETIME);`
	query, err := db.Prepare(users_table)
	if err != nil {
		log.Fatal(err)
	}
	query.Exec()
	fmt.Println("Table created successfully!")
}

func DatabaseInit() {
	CreateDatabase()
	db := DBConnect()
	CreateBitcoinTrackerTable(db)
}

func DBConnect() *sql.DB {
	// Connect to database
	db, err := sql.Open("sqlite3", "./database.CryptoTracker")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
