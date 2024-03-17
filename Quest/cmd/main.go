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

func main() {
	/*router := gin.Default()
	router.GET("/", func(c *gin.Context) {

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses (in this case, 'title')
			gin.H{
				"title": "Home Page",
			},
		)

	})
	router.Run()*/
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
