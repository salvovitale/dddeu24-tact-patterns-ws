package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/domain"
	infra_repository "github.com/salvovitale/dddeu24-tact-patterns-ws/internal/infra/repository"
	"github.com/salvovitale/dddeu24-tact-patterns-ws/internal/web"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	priceSvc := domain.NewPriceSvc(&infra_repository.UserRepository{})
	h := web.NewHandler(logger, priceSvc)
	slog.Info("starting server", slog.Any("port", 5000))
	if err := http.ListenAndServe(":5000", h); err != nil {
		slog.Error("error starting server", slog.Any("error", err.Error()))
		panic(err)
	}
}
