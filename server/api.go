package server

import (
	"net/http"
	"strconv"

	"github.com/urfave/negroni"
	"gopkg.in/tylerb/graceful.v1"
	"nipun.io/message_queue/appcontext"
	"nipun.io/message_queue/config"
	"nipun.io/message_queue/logger"
	r "nipun.io/message_queue/server/router"
)

func listenServer(apiServer *graceful.Server) {
	logger.Logger.Info().Msgf("starting api server on address : %s", apiServer.Addr)
	if err := apiServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Logger.Fatal().Err(err)
	}
}

func StartApiServer(dependencies *appcontext.Instance) {
	logger.Logger.Info().Msg("Starting API server")

	router := r.Router(dependencies)
	n := negroni.New(negroni.NewRecovery())
	n.UseHandlerFunc(router.ServeHTTP)

	portInfo := ":" + strconv.Itoa(config.AppPort())
	n.Run(portInfo)
}
