package taskboard

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

// InitMySQLClient initializes the MySQL client at the start of the process.
func init() {
	dsn := "exampleuser:examplepassword@tcp(localhost:3306)/exampledb"

	once.Do(func() {
		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Failed to connect to MySQL: %v", err)
		}

		// Test the connection
		if err = db.Ping(); err != nil {
			log.Fatalf("Failed to ping MySQL: %v", err)
		}

		fmt.Println("MySQL client initialized successfully")
	})
}

// GetMySQLClient exposes the MySQL client instance.
func GetMySQLClient() *sql.DB {
	if db == nil {
		log.Fatal("MySQL client is not initialized. Call InitMySQLClient first.")
	}
	return db
}
