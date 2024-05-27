package web

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
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
		type DroppedFraction struct {
			AmountDropped float64 `json:"amount_dropped"`
			FractionType  string  `json:"fraction_type"`
		}
		type Request struct {
			Date             string            `json:"date"`
			DroppedFractions []DroppedFraction `json:"dropped_fractions"`
			PersonID         string            `json:"person_id"`
			VisitID          string            `json:"visit_id"`
		}

		var req Request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var fractions []domain.Fraction

		for _, f := range req.DroppedFractions {
			ft, err := domain.ParseFractionType(f.FractionType)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			weight, err := domain.ParseWeight(f.AmountDropped)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fractions = append(fractions, domain.Fraction{
				Type: ft,
				Kg:   weight,
			})
		}

		p := domain.NewPrice(fractions)
		price := p.CalculatePrice()

		w.Header().Set("Content-Type", "application/json")

		type response struct {
			PriceAmount   float64 `json:"price_amount"`
			PersonID      string  `json:"person_id"`
			VisitID       string  `json:"visit_id"`
			PriceCurrency string  `json:"price_currency"`
		}
		responseData := response{
			PriceAmount:   price,
			PersonID:      req.PersonID,
			VisitID:       req.VisitID,
			PriceCurrency: "EUR",
		}
		json.NewEncoder(w).Encode(responseData)
	}
}
