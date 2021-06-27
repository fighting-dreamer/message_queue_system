package domain

import (
	"encoding/json"
	"net/http"

	"nipun.io/message_queue/logger"
)

type errorResponse struct {
	Errors []string `json:"errors"`
}

func WriteResponse(status int, response interface{}, rw http.ResponseWriter) {
	if response == nil {
		response = struct{}{}
	}
	body, err := json.Marshal(response)
	if err != nil {
		logger.Logger.Error().Err(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(body)
}

func WriteErrorResponse(status int, errorMessages []string, rw http.ResponseWriter) {
	if errorMessages == nil {
		errorMessages = []string{}
	}
	errorResponse := errorResponse{Errors: errorMessages}
	respBytes, err := json.Marshal(errorResponse)
	if err != nil {
		logger.Logger.Error().Err(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
