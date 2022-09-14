package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	url := "https://klike.net/uploads/posts/2019-05/1556708064_4.jpg"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("can't get resp, %v\n", err)
		return
	}

	defer func() {
		if e := resp.Body.Close(); e != nil {
			log.Fatalf("failed to close file %v, error: %q", resp, e)
			return
		}
	}()

	body, err := io.ReadAll(resp.Body)

	contentType := Sniff(body)

	if strings.Contains(contentType, "image") {
		fmt.Printf("Это изображение")
	} else {
		fmt.Printf("Это НЕ изображение")
	}
}
