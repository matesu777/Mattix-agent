package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/matesu777/Mattix/internal/handle"

	"github.com/matesu777/Mattix/internal/collector"
)

func main() {
	col, err := collector.New()
	if err != nil {
		log.Fatalf("Fatal error to initialize application: %v", err)
	}

	go col.SlowStart()
	go col.FastStart()

	handler := handle.New(col)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /metrics", handler.GetMetrics)

	fmt.Printf("Mattix agent v0.1.0 \n\nhostname: %s \nlistening on: 8080\n", col.Metrics.Hostname)

	log.Fatal(http.ListenAndServe(":8080", handle.Cors(mux)))
}
