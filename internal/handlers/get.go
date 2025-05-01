package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"path"

	"github.com/ensomnatt/ducks/internal/logger"
	"github.com/jackc/pgx/v5"
)

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	name := path.Base(r.URL.Path)

	logger.Log.Debug("got request", "request", "get a duck", "name", name)

	ctx, cancel := h.CreateContext()
	defer cancel()

	duck, err := h.db.Get(name, ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Log.Info("bad duck's name", "name", name)
			http.Error(w, "bad duck's name", http.StatusBadRequest)
			return
		} else {
			h.HandleSendingError(w, err)
			return
		}
	}

	logger.Log.Info("got a duck", "name", duck.Name, "age", duck.Age)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	logger.Log.Debug("setted up a header")

	err = json.NewEncoder(w).Encode(duck)
	if err != nil {
		return
	}

	logger.Log.Debug("sent a response")
	return
}
