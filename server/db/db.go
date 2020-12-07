package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

// DB refers to the connected db.
var DB *pg.DB

type InitDBOpts struct {
	User     string
	Database string
}

const (
	defDBUser   = "kunalpowar"
	defDatabase = "gopggqlauth"
)

func (o *InitDBOpts) useDefaultsIfEmpty() {
	if o.User == "" {
		o.User = defDBUser
	}
	if o.Database == "" {
		o.Database = defDBUser
	}
}

// Init connects to the db and initialises the global DB var.
func Init(opts *InitDBOpts) {
	opts.useDefaultsIfEmpty()

	DB = pg.Connect(&pg.Options{
		User:     opts.User,
		Database: opts.Database,
	})

	mustPing()

	log.Printf("connected to DB: %v", DB)
}

func mustPing() {
	if err := DB.Ping(context.Background()); err != nil {
		log.Fatalf("db: could not ping db: %v", err)
	}
}
