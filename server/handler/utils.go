package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"nipun.io/message_queue/logger"
)

func getBody(ctx context.Context, reqBodyByte []byte, body interface{}) error {
	err := json.Unmarshal(reqBodyByte, body)
	if err != nil {
		status := http.StatusBadRequest
		logMessage := "Failed parsing request payload for returning %d, err: %v"
		logger.Logger.Error().Msg(fmt.Sprintf(logMessage, status, err))
	}
	return err
}
