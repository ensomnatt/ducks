package main

import (
	"github.com/ensomnatt/ducks/internal/db"
	"github.com/ensomnatt/ducks/internal/handlers"
	"github.com/ensomnatt/ducks/internal/logger"
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

	handlers.Start(db)
}
