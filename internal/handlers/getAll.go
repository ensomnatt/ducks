package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ensomnatt/ducks/internal/logger"
)

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	logger.Log.Debug("got request", "request", "get all ducks")

	ctx, cancel := h.CreateContext()
	defer cancel()

	ducks, err := h.db.GetAll(ctx)
	if err != nil {
		logger.Log.Error("error while getting all ducks", "error", err)
		http.Error(w, "failed to get all ducks", http.StatusInternalServerError)
		return
	}

	logger.Log.Debug("got all ducks", "ducks", ducks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	logger.Log.Debug("setted up a header")

	err = json.NewEncoder(w).Encode(&ducks)
	if err != nil {
		h.HandleSendingError(w, err)
		return
	}

	logger.Log.Info("got all ducks", "ducks", ducks)
	return
}
