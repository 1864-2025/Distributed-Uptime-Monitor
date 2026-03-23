package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func connectDB() *pgxpool.Pool {
	dsn := "postgres://user:password@localhost:5432/uptime_monitor"

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to database")

	return dbpool
}
