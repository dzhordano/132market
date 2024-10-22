package app

import (
	"dzhordano/132market/services/gateway/internal/config"
	"dzhordano/132market/services/gateway/internal/delivery/http"
	"dzhordano/132market/services/gateway/internal/server"
	"log"
)

func Run() {
	h := http.New()

	s := server.New(config.NewHttpConfig().Address(), h.InitRoutes())

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
