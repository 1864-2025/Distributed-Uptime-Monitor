package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func loadUrlsFromJson(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var urls []string
	err = json.Unmarshal(file, &urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

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

	urls, err := loadUrlsFromJson("urls.json")
	if err != nil {
		log.Fatalf("Не удалось загрузить список URL: %v", err)
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
