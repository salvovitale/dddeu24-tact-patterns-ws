package web

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	slogchi "github.com/samber/slog-chi"
)

func NewHandler(logger *slog.Logger) *Handler {

	h := &Handler{
		Mux: chi.NewRouter(),
	}

	// add logger middleware
	h.Use(slogchi.New(logger))

	// root path
	h.Get("/", h.statusOK())
	h.Post("/startScenario", h.startScenario())
	h.Post("/calculatePrice", h.calculatePrice())

	// sub paths
	// h.Route("/threads", func(r chi.Router) {
	// 	r.Get("/", threadsHandler.listView())
	// })
	return h
}

type Handler struct {
	*chi.Mux //embedded structure
}

func (h *Handler) statusOK() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonData := map[string]string{"status": "OK"}
		json.NewEncoder(w).Encode(jsonData)
	}
}

func (h *Handler) startScenario() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		jsonData := map[string]string{}
		json.NewEncoder(w).Encode(jsonData)
	}
}

func (h *Handler) calculatePrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		type response struct {
			PriceAmount   float64 `json:"price_amount"`
			PersonID      string  `json:"person_id"`
			VisitID       string  `json:"visit_id"`
			PriceCurrency string  `json:"price_currency"`
		}
		responseData := response{
			PriceAmount:   0,
			PersonID:      "Bald Eagle",
			VisitID:       "1",
			PriceCurrency: "EUR",
		}
		json.NewEncoder(w).Encode(responseData)
	}
}
