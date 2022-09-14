package main

import (
	"fmt"
	"github.com/aofei/mimesniffer"
	"io"
	"net/http"
)

func main() {

	url := "https://klike.net/uploads/posts/2019-05/1556708064_4.jpg"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("can't get resp, %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	contentType := mimesniffer.Sniff(body)

	fmt.Printf("Тип файла: %v\n", contentType)

}
