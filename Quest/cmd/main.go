package main

import (
	"Quest/internal/handler"
	"Quest/internal/postgres"
	"Quest/internal/server"
	"Quest/internal/service"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Your Gin Application
// @version 1.0
// @description This is a sample server for your Gin Application.
// @BasePath /
func main() {
	err := godotenv.Load()

	conn, err := postgres.InitConn()
	if err != nil {
		log.Println(err)
		log.Fatal("Can't init connection to database")
	}
	defer conn.Close()

	db := postgres.InitDb(conn)

	service := service.InitService(db)

	handler := handler.InitHandler(service)

	server.Run(handler)

	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	<-c
}
