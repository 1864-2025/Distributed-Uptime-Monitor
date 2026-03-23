package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func connectDB() *pgx.Conn {
	dsn := "postgres://user:password@localhost:5432/uptime_monitor"

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to database")

	return conn
}
