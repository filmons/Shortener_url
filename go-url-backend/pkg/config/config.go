// package config

// import (
// 	"fmt"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// //ConnectDB connects go to mysql database
// func ConnectDB() *gorm.DB {
// 	errorENV := godotenv.Load()
// 	if errorENV != nil {
// 		panic("Failed to load env file")
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbName := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
// 	db, errorDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if errorDB != nil {
// 		panic("Failed to connect mysql database")
// 	}

// 	return db
// }

// //DisconnectDB is stopping your connection to mysql database
// func DisconnectDB(db *gorm.DB) {
// 	dbSQL, err := db.DB()
// 	if err != nil {
// 		panic("Failed to kill connection from database")
// 	}
// 	dbSQL.Close()
// }
package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ConnectDB connects to the MySQL database.
func ConnectDB() *gorm.DB {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load env file: %v", err)
	}

	// Retrieve database connection parameters from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL database: %v", err)
	}

	log.Println("Connected to MySQL database")

	return db
}

// DisconnectDB disconnects from the MySQL database.
func DisconnectDB(db *gorm.DB) {
	// Close the database connection
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}
	if err := dbSQL.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}

	log.Println("Disconnected from MySQL database")
}
