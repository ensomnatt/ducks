package main

import (
	"context"
	"time"

	"github.com/ensomnatt/ducks/internal/db"
	"github.com/ensomnatt/ducks/internal/logger"
	"github.com/ensomnatt/ducks/internal/models"
)

func main() {
	//TODO:
	//сделать хендлеры

	conn, err := db.ConnectToDB()
	if err != nil {
		logger.Log.Error("failed to connect to db", "error", err)
		return
	}
	db := db.NewDucksDB(conn)
	db.Init()
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.Create(models.Duck{Name: "zalupa", Age: 15}, ctx)
	logger.Log.Info("created a duck")
	
	err = db.Create(models.Duck{Name: "pidor", Age: 18}, ctx)
	logger.Log.Info("created a duck")

	duck, err := db.Get("zalupa", ctx)
	logger.Log.Info("got a duck", "name", duck.Name, "age", duck.Age)
	
	ducks, err := db.GetAll(ctx)
	logger.Log.Info("got ducks", "ducks", ducks)

	if err != nil {
		logger.Log.Error("error", "error", err)
		return
	}
}
