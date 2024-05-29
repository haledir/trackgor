package main

import (
	"context"
	"log"
	"os"

	"github.com/haledir/trackgor/db"
	"github.com/haledir/trackgor/views"
)

func main() {
	database, err := db.InitDB("./trackgor.db")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer database.Close()

	component := views.Hello("Michael")
	component.Render(context.Background(), os.Stdout)

}
