package main

import (
	"encoding/json"
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

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(collector.Metrics)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
