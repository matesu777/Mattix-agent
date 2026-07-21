package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/matesu777/Mattix/internal/collector"
)

func main() {
	collector, err := collector.New()
	if err != nil {
		log.Fatal(err)
	}

	collector.Update()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			if err := collector.Update(); err != nil {
				log.Println(err)
			}
		}
	}()

	fmt.Printf("Mattix agent v0.1.0 \n\nhostname: %s \nlistening on: 8080\n", collector.Metrics.Hostname)

	mux := http.NewServeMux()

	mux.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(collector.Metrics)
	})

	log.Fatal(http.ListenAndServe(":8080", cors(mux)))
}

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
