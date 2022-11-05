package main

import (
	"context"
	"os"
	"os/signal"
	"quizbackend/internal/api"
	"quizbackend/internal/utils/db"
	"quizbackend/internal/utils/logger"
	"time"
)

func main() {
	//start database

	db.StartDB()
	defer db.Db.Close()

	//start server
	server := api.StartServer()

	go func() {
		logger.InfoLogger.Println("server: Server started on port 8080")
		if err := server.ListenAndServe(); err != nil {
			logger.FatalLogger.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1) //listens for cancellation
	signal.Notify(c, os.Interrupt)
	<-c //pause code here until interrupted
	logger.InfoLogger.Println("server: Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	server.Shutdown(ctx)
}
