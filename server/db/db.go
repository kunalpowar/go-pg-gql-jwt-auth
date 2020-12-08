package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

// Instance refers to the connected db.
var Instance *pg.DB

// MustInit connects to the db and initialises the global DB var.
func MustInit(ctx context.Context) {
	opts := pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		Database: os.Getenv("DB_NAME"),
	}

	tcpHost := os.Getenv("DB_TCP_HOST")
	if tcpHost != "" {
		opts.Network = "tcp"
		opts.Addr = tcpHost

		log.Printf("db: attempting to connect to db via tcp: %s", opts.Addr)
	} else {
		dbSocketName := os.Getenv("INSTANCE_CONNECTION_NAME")
		if dbSocketName == "" {
			log.Fatalf("db: at least one of DB_TCP_HOST or INSTANCE_CONNECTION_NAME should be set in env")
		}

		opts.Network = "unix"
		opts.Addr = fmt.Sprintf("/cloudsql/%s/.s.PGSQL.5432", dbSocketName)

		log.Printf("db: attempting to connect to db via unit socket: %s", opts.Addr)
	}

	Instance = pg.Connect(&opts)

	if err := Instance.Ping(ctx); err != nil {
		log.Fatalf("db: could not ping db: %v", err)
	}

	log.Printf("db: connection established: %v", Instance)
}
