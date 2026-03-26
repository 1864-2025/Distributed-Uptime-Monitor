package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func checkURL(url string, wg *sync.WaitGroup, conn *pgxpool.Pool) {

	defer wg.Done()

	start := time.Now()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)

	duration := time.Since(start).Milliseconds()

	statusCode := 0

	if err != nil {
		fmt.Printf("Ошибка при проверке %s: %v\n", url, err)
	} else {
		statusCode = resp.StatusCode
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Printf("Error closing body: %v\n", err)
			}
		}(resp.Body)
	}

	query := `INSERT INTO checks (url, status_code, response_time_ms) VALUES ($1, $2, $3)`

	_, err = conn.Exec(context.Background(), query, url, statusCode, duration)
	if err != nil {
		fmt.Printf("Ошибка записи в базу: %v\n", err)
	}

	fmt.Printf("URL: %s; Status: %d; Time: %dms\n", url, statusCode, duration)

	return
}

func main() {

	conn := connectDB()
	defer func(p *pgxpool.Pool) {
		p.Close()
		fmt.Println("Database connection pool closed.")
	}(conn)

	wg := sync.WaitGroup{}

	urls := []string{
		"https://www.google.com",
		"https://www.yandex.ru",
		"https://www.cu.ru",
		"https://www.facebook.com",
		"https://www.twitter.com",
		"https://www.youtube.com",
		"https://www.golgfhisuhefol.com",
	}

	for {
		wg.Add(len(urls))
		for _, url := range urls {
			go checkURL(url, &wg, conn)
		}

		wg.Wait()

		fmt.Println("Iteration end. Waiting 60 seconds.")

		time.Sleep(60 * time.Second)

	}
}
