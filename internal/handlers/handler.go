package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/ensomnatt/ducks/internal/db"
	"github.com/ensomnatt/ducks/internal/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	db *db.DucksDB
}

func Start(db *db.DucksDB) {
	r := http.NewServeMux()
	h := Handler{
		db: db,
	}

	r.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("CREATE /api/create", h.Create)
	r.HandleFunc("GET /api/{name}", h.Get)
	r.HandleFunc("GET /api/all", h.GetAll)

	http.ListenAndServe(":4242", r)
}

func (h Handler) HandleSendingError(w http.ResponseWriter, err error) {
	logger.Log.Error("error while getting a duck", "error", err)
	http.Error(w, "failed to get a duck", http.StatusInternalServerError)
}

func (h Handler) CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
