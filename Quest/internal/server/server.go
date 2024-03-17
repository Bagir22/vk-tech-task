package server

import (
	"Quest/internal/handler"
)

func Run(h *handler.Handler) {
	h.Init()
	router := h.Init()
	router.Run()
}
