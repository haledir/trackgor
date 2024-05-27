package main

import (
	"log"

	"github.com/haledir/trackgor/db"
)

func main() {
	database, err := db.InitDB("./trackgor.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()

}
