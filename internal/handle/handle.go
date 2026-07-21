package handle

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matesu777/Mattix/internal/collector"
)

type Handler struct {
	collector *collector.Collector
}

func New(c *collector.Collector) *Handler {
	return &Handler{
		collector: c,
	}
}

func (h *Handler) GetMetrics(w http.ResponseWriter, r *http.Request) {
	metrics := h.collector.GetMetrics()

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(metrics); err != nil {
		log.Println(err)
	}
}
