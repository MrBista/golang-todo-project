package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/MrBista/golang-todo-project/helper"
	"github.com/MrBista/golang-todo-project/src/dto/response"
	"github.com/julienschmidt/httprouter"
)

func AutthMiddlware(handle httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		bearerToken := r.Header.Get("Authorization")
		if bearerToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := response.CommonResponse{
				Data:    false,
				Status:  http.StatusUnauthorized,
				Message: "unauthorized user - missing token",
			}
			json.NewEncoder(w).Encode(webResponse)
			return
		}
		tokenParts := strings.Split(bearerToken, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := response.CommonResponse{
				Data:    false,
				Status:  http.StatusUnauthorized,
				Message: "invalid token format",
			}
			json.NewEncoder(w).Encode(webResponse)
			return
		}

		accessToken := tokenParts[1]

		err := helper.VerifyToken(accessToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			webResponse := response.CommonResponse{
				Data:    false,
				Status:  http.StatusUnauthorized,
				Message: "invalid or expired token",
			}
			json.NewEncoder(w).Encode(webResponse)
			return
		}

		// Token is valid, proceed to original handler
		handle(w, r, p)
	}
}
