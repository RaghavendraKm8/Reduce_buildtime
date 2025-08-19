package main

import (
	"fmt"
	"net/http"

	"myapp/internal/compute"

	"github.com/google/uuid"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n := 25
		fib := compute.Fib(n)
		id := uuid.New() // external dependency
		fmt.Fprintf(w, "hello from myapp! fib(%d)=%d uuid=%s\n", n, fib, id)
	})

}
