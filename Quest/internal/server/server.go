package server

import (
	"Quest/internal/config"
	"Quest/internal/handler"
	"fmt"
	"log"
	"net/http"
)

func Run(h *handler.Handler) {
	cfg := config.InitConfig()

	s := &http.Server{
		Addr:    fmt.Sprintf("localhost:%v", cfg.ServerPort),
		Handler: h.Init(),
	}

	log.Printf("Run server at localhost:%v\n", cfg.ServerPort)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatalf("HTTP server ListenAndServe Error: %v", err)
		}
	}()
}
