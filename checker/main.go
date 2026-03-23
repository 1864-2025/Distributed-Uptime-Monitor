package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func checkURL(url string, wg *sync.WaitGroup) {

	defer wg.Done()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Ошибка при проверке %s: %v\n", url, err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing body: %v\n", err)
		}
	}(resp.Body)

	fmt.Printf("URL: %s; response Status: %s\n", url, resp.Status)

	return
}

func main() {
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
	wg.Add(len(urls))
	for _, url := range urls {
		go checkURL(url, &wg)
	}
	wg.Wait()
}
