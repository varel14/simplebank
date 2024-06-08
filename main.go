package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	db "github.com/varelx/bank/db/sqlc"
)

const (
	connString = "postgres://varel:newpassword237@localhost:5432/super_bank?sslmode=disable"
)

func main() {
	conn, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		log.Fatal("enable to connect to db")
	}
	q := db.New(conn)

	fmt.Println("db connected", q)
}
