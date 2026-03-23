package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	url := "https://www.google.com"

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

		}
	}(resp.Body)

	fmt.Println("response Status:", resp.Status)
}
