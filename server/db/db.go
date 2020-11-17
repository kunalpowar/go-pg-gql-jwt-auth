package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

// DB refers to the connected db.
var DB *pg.DB

// Init connects to the db and initialises the global DB var.
func Init() {
	DB = pg.Connect(&pg.Options{
		User:     "kunalpowar",
		Database: "gopggqlauth",
	})

	mustPing()

	log.Printf("connected to DB: %v", DB)
}

func mustPing() {
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalf("db: could not ping db: %v", err)
	}
}
