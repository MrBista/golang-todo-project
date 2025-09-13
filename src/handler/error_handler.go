package handler

import (
	"encoding/json"
	"net/http"

	"github.com/MrBista/golang-todo-project/src/exception"
)

type CommoneErrorResponse struct {
	Message string            `json:"message"`
	Code    string            `json:"code"`
	Errors  map[string]string `json:"errors,omitempty"`
	Status  int               `json:"status"`
}

func HandleError(w http.ResponseWriter, err error) {

	if custmError, ok := exception.IsCustomError(err); ok {
		writeErrorCustom(w, custmError)
		return
	}

	internalServerError := exception.NewInternalServerError("Internal server error")
	writeErrorCustom(w, internalServerError)
}

func writeErrorCustom(w http.ResponseWriter, custmError *exception.ErrorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(custmError.GetStatusHttp())
	errResponse := CommoneErrorResponse{
		Message: custmError.Message,
		Code:    custmError.Code,
		Errors:  custmError.Errors,
		Status:  custmError.GetStatusHttp(),
	}

	encoder := json.NewEncoder(w)

	err := encoder.Encode(errResponse)
	if err != nil {
		panic("Failed write response")
	}
}
