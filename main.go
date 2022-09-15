package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {

	url := "https://static.eldorado.ru/photos/mv/Big/30058843bb1.jpg"

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

	body := make([]byte, 512)
	body, err = io.ReadAll(resp.Body)

	contentType := Sniff(body)

	if strings.Contains(contentType, "image") {
		fmt.Printf("This is image")
	} else {
		fmt.Printf("This is NOT image")
	}
}
