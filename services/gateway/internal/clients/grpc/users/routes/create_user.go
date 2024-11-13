package routes

import (
	"context"
	"dzhordano/132market/services/users/pkg/pb/user_v1"
	"encoding/json"
	"log"
	"net/http"
)

type CreateUserReqBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(w http.ResponseWriter, r *http.Request, c user_v1.UserServiceV1Client) {
	reqBody := &CreateUserReqBody{}
	if err := json.NewDecoder(r.Body).Decode(reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	command := &user_v1.CreateUserRequest{
		Info: &user_v1.UserInfo{
			Name:     reqBody.Name,
			Email:    reqBody.Email,
			Password: reqBody.Password,
		},
	}

	res, err := c.CreateUser(context.TODO(), command)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}
