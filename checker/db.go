package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func connectDB() *pgxpool.Pool {

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("Предупреждение: .env файл не найден, используются системные переменные")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbName)

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		fmt.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully connected to database")

	return dbpool
}
