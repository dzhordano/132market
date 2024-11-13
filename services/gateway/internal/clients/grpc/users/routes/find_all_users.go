package routes

import (
	"context"
	"dzhordano/132market/services/users/pkg/pb/user_v1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// FIXME РЕФАКТОР

func validateIntParam(param string) (uint64, error) {
	if param == "" {
		return 1, nil
	}

	paramInt, err := strconv.ParseUint(param, 10, 64)
	if err != nil {
		return 0, err
	}

	if param == "limit" && (paramInt < 1 || paramInt > 100) {
		return 0, fmt.Errorf("limit param must be between 1 and 100")
	}

	return uint64(paramInt), nil
}

func FindAllUsers(w http.ResponseWriter, r *http.Request, c user_v1.UserServiceV1Client) {
	var limitInt uint64
	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" {
		limitInt = 10
	} else {
		var err error
		limitInt, err = validateIntParam(limitParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	var offsetInt uint64
	offsetParam := r.URL.Query().Get("offset")
	if offsetParam == "" {
		offsetInt = 0
	} else {
		var err error
		offsetInt, err = validateIntParam(offsetParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)

			fmt.Println("ERR", err)

			json.NewEncoder(w).Encode(err)
			return
		}
	}

	res, err := c.FindAllUsers(context.TODO(), &user_v1.FindAllUsersRequest{
		Offset: offsetInt,
		Limit:  limitInt,
	})
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
