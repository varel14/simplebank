package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	db "github.com/varelx/bank/db/sqlc"
)

const (
	connString = "postgres://varel:newpassword237@localhost:5432/super_bank?sslmode=disable"
)

func main() {
	// var testQueries *Queries
	// var testDB *pgx.Conn
	var ctx = context.Background()

	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
		log.Fatal("enable to connect to db")
	}
	q := db.New(conn)
	fmt.Println("hello world")

	store := db.NewStore(conn)

	accounts, err := q.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  10,
		Offset: 0,
	})
	if err != nil || len(accounts) < 2 {
		log.Fatal("less than 02 accounts")
	}

	result, err := store.TransferTx(context.Background(), db.TransferTxParams{
		FromAccountID: accounts[0].ID,
		ToAccountID:   accounts[1].ID,
		Amount:        120,
	})

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ok", result)
}
