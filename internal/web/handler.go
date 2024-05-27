package web

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
	slogchi "github.com/samber/slog-chi"
)

type Handler struct {
	*chi.Mux //embedded structure
	priceSvc domain.PriceService
}

func NewHandler(logger *slog.Logger, priceSvc domain.PriceService) *Handler {

	h := &Handler{
		Mux:      chi.NewRouter(),
		priceSvc: priceSvc,
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

		var req CalculatePriceRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			slog.Error("error parsing weight", slog.Any("error", err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var fractions []domain.Fraction

		for _, f := range req.DroppedFractions {
			ft, err := domain.ParseFractionType(f.FractionType)
			if err != nil {
				slog.Error("error parsing weight", slog.Any("error", err))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			weight, err := domain.ParseWeight(f.AmountDropped)
			if err != nil {
				slog.Error("error parsing weight", slog.Any("error", err))
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fractions = append(fractions, domain.Fraction{
				Type: ft,
				Kg:   weight,
			})
		}
		price, err := h.priceSvc.CalculatePrice(fractions, req.PersonID)
		if err != nil {
			slog.Error("error calculate price", slog.Any("error", err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		responseData := CalculatePriceResponse{
			PriceAmount:   price,
			PersonID:      req.PersonID,
			VisitID:       req.VisitID,
			PriceCurrency: "EUR",
		}
		json.NewEncoder(w).Encode(responseData)
	}
}
