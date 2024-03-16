package main

import (
	"Quest/internal/handler"
	"Quest/internal/postgres"
	"Quest/internal/server"
	"Quest/internal/service"
	"github.com/joho/godotenv"
	"log"
)

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
}
