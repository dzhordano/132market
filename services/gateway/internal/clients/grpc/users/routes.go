package users

import (
	"context"
	"dzhordano/132market/services/gateway/internal/clients/grpc/users/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InitRoutes(r chi.Router) *UserClient {

	svc := &UserClient{
		c: NewUserClient(context.TODO()),
	}

	r.Post("/users/create", svc.CreateUser)
	r.Get("/users/get", svc.FindAllUsers)

	return svc
}

func (uc *UserClient) CreateUser(w http.ResponseWriter, r *http.Request) {
	routes.CreateUser(w, r, uc.c)
}

func (uc *UserClient) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	routes.FindAllUsers(w, r, uc.c)
}
