package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

// Instance refers to the connected db.
var Instance *pg.DB

type InitDBOpts pg.Options

const (
	defAddres   = "localhost:5432"
	defDBUser   = "kunalpowar"
	defDatabase = "gopggqlauth"
)

func (o *InitDBOpts) useDefaultsIfEmpty() {
	if o.Addr == "" {
		o.Addr = defAddres
	}

	if o.User == "" {
		o.User = defDBUser
	}

	if o.Database == "" {
		o.Database = defDBUser
	}
}

// MustInit connects to the db and initialises the global DB var.
func MustInit(opts *InitDBOpts) {
	opts.useDefaultsIfEmpty()

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
