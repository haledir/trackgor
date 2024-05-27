package db

import (
	"database/sql"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func InitDB(dbName string) (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, err
	}

	migration := `create table if not exists ARTICLES (
        ID integer primary key autoincrement,
        TITLE text not null,
        CONTENT text not null
    );`
	_, err = db.Exec(migration)
	if err != nil {
		return nil, err
	}

	migration_user := `create table if not exists USERS (
        ID integer primary key autoincrement,
        USERNAME text unique not null,
        PASSWORD text not null
    );`
	_, err = db.Exec(migration_user)
	if err != nil {
		return nil, err
	}
	insertInitialCredentials(db)
	return db, nil
}

func insertInitialCredentials(db *sql.DB) {
	username := os.Getenv("INITIAL_USER")
	password := os.Getenv("INITIAL_PASSWORD")
	if username == "" || password == "" {
		log.Println("Initial credentials not set in .env file")
		return
	}

	var id int

	err := db.QueryRow("SELECT id FROM users WHERE username = ?", username).Scan(&id)
	if err == sql.ErrNoRows {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
		if err != nil {
			log.Fatalf("Error hashing password: %v", err)
		}
		_, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
		if err != nil {
			log.Fatalf("Error inserting initial user: %v", err)
		}
		log.Println("Initial user created with username: test and password: test")
	} else if err != nil {
		log.Fatalf("Error checking initial user: %v", err)
	}
}
