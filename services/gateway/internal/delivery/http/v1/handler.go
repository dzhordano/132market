package v1

import (
	"dzhordano/132market/services/gateway/internal/clients/grpc/users"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handlerV1 struct {
}

func New() *handlerV1 {
	return &handlerV1{}
}

func (h *handlerV1) InitRoutes() chi.Router {

	r := chi.NewRouter()

	r.Route("/v1", func(r chi.Router) {
		users.InitRoutes(r)

		r.Mount("/auth", initSSORoutes())
		//r.Mount("/users", initUsersRoutes())
		r.Mount("/products", initProductsRoutes())
		r.Mount("/orders", initOrdersRoutes())
		r.Mount("/statistics", initStatisticsRoutes())
		r.Mount("/chat", initChatRoutes())
		r.Mount("/mailing", initMailingRoutes())
		r.Mount("/admin", initAdminRoutes())
	})

	return r
}

// Dummy temporary method.
func notImplemented(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
