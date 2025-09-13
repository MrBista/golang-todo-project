package exception

import (
	"encoding/json"
	"net/http"

	"github.com/MrBista/golang-todo-project/src/dto/response"
)

func ErrorHandler(w http.ResponseWriter, req *http.Request, err interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := response.CommonResponse{
		Data:    false,
		Status:  http.StatusInternalServerError,
		Message: "INTERNAL_SERVER_ERROR",
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(webResponse)
}
