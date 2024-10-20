package app

import (
	"dzhordano/132market/services/gateway/internal/delivery/http"
	"log"
	httpServ "net/http"
)

func Run() {
	h := http.New()

	err := httpServ.ListenAndServe(":8080", h.InitRoutes())
	if err != nil {
		log.Fatal(err)
	}
}
