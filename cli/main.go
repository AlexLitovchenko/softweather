package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("Need string and URL in argmuments")
	}

	str := os.Args[1]
	url := os.Args[2]

	resp, err := http.Post(url, "text/plain", strings.NewReader(str))
	if err != nil {
		log.Fatal("Failed with post request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response:", err)
	}

	fmt.Println(string(body))
}
