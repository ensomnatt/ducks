package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ensomnatt/ducks/internal/logger"
	"github.com/ensomnatt/ducks/internal/metrics"
	"github.com/ensomnatt/ducks/internal/models"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var req models.Duck
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logger.Log.Error("error while decoding request body", "error", err)
		http.Error(w, "failed to get your request", http.StatusInternalServerError)
		return
	}

	logger.Log.Debug("got request", "request", "create a duck", "name", req.Name, "age", req.Age)

	ctx, cancel := h.CreateContext()
	defer cancel()

	err = h.db.Create(req, ctx)

	if err != nil {
		logger.Log.Error("error while creating a duck", "error", err)
		http.Error(w, "failed to create a duck", http.StatusInternalServerError)
		return
	}

	logger.Log.Info("created a duck", "name", req.Name, "age", req.Age)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	logger.Log.Debug("setted up a header")

	duration := time.Since(start).Seconds()

	metrics.HttpRequests.WithLabelValues(
		r.Method,
		r.URL.Path,
		strconv.Itoa(http.StatusCreated),
	).Inc()
	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	logger.Log.Debug("updated metrics")

	return
}
