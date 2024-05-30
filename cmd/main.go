package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/haledir/trackgor/db"
	"github.com/haledir/trackgor/views"
)

func main() {
	database, err := db.InitDB("./trackgor.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()

	users, err := db.GetUsers(database)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	indexComponent := views.Index(users)

	http.Handle("/", templ.Handler(indexComponent))
	http.ListenAndServe(":42069", nil)

}
