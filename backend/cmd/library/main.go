package main

import (
	"fmt"
	"library/pkg/conf"
	"library/pkg/server"
	"net/http"

	"github.com/rs/zerolog/log"
)

func main() {

	config := conf.NewConfig("dev")

	s := server.NewServer(config)

	s.InitAuthorRoutes()

	log.Error().Err(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port), s.Router))

}
