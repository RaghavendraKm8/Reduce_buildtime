package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback
	}

	url := fmt.Sprintf("http://localhost:%s", port)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("❌ Failed to reach server:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("✅ Server response:\n", string(body))
}
