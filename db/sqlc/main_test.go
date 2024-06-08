package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	connString = "postgres://varel:newpassword237@localhost:5432/super_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool
var ctxx = context.Background()

func TestMain(m *testing.M) {
	var err error

	testDB, err = pgxpool.New(ctxx, connString)
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
