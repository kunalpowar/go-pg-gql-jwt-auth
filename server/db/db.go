package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

// Instance refers to the connected db.
var Instance *pg.DB

type InitDBOpts pg.Options

// MustInit connects to the db and initialises the global DB var.
func MustInit(opts *InitDBOpts) {
	Instance = pg.Connect(&pg.Options{
		User:     opts.User,
		Database: opts.Database,
	})

	mustPing()

	log.Printf("connected to DB: %v", Instance)
}

func mustPing() {
	if err := Instance.Ping(context.Background()); err != nil {
		log.Fatalf("db: could not ping db: %v", err)
	}
}
