package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/ensomnatt/ducks/internal/db"
	"github.com/ensomnatt/ducks/internal/logger"
)

type Handler struct {
	db *db.DucksDB
}

func Start(db *db.DucksDB) {
	r := http.NewServeMux()
	h := Handler{
		db: db,
	}

	r.HandleFunc("CREATE /", h.Create)
	r.HandleFunc("GET /{name}", h.Get)
	r.HandleFunc("GET /", h.GetAll)
}

func (h Handler) HandleSendingError(w http.ResponseWriter, err error) {
	logger.Log.Error("error while getting a duck", "error", err)
	http.Error(w, "failed to get a duck", http.StatusInternalServerError)
}

func (h Handler) CreateContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}
