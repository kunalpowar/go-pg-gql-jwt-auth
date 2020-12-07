package db

import (
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

	log.Printf("connected to DB: %v", Instance)
}
