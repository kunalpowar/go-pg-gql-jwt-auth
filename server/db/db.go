package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	// Postgres driver
	_ "github.com/jackc/pgx/stdlib"
)

// Instance refers to the connected db.
var Instance *sqlx.DB

// MustInit connects to the db and initialises the global DB var.
func MustInit(ctx context.Context) {
	var (
		user         = os.Getenv("DB_USER")
		password     = os.Getenv("DB_PASS")
		database     = os.Getenv("DB_NAME")
		tcpHost      = os.Getenv("DB_TCP_HOST")
		tcpPort      = os.Getenv("DB_TCP_PORT")
		dbSocketName = os.Getenv("INSTANCE_CONNECTION_NAME")

		connectionString = fmt.Sprintf("dbname=%s", database)
	)

	if tcpHost != "" {
		connectionString = fmt.Sprintf("host=%s port=%s %s", tcpHost, tcpPort, connectionString)
		log.Printf("db: attempting to connect to db via tcp: %q", connectionString)
	} else if dbSocketName != "" {
		connectionString = fmt.Sprintf("host=/cloudsql/%s/.s.PGSQL.5432 %s user=%s password=%s", dbSocketName, connectionString, user, password)
		log.Printf("db: attempting to connect to db via unit socket: %q", connectionString)
	} else {
		log.Fatalf("db: at least one of DB_TCP_HOST or INSTANCE_CONNECTION_NAME should be set in env")
	}

	var err error
	Instance, err = sqlx.Connect("pgx", connectionString)
	if err != nil {
		log.Fatalf("db: could not connect to db: %v", err)
	}

	log.Printf("db: connection established: %v", Instance)
}
