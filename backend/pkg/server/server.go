package server

import (
	"context"
	"fmt"
	"library/pkg/conf"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type Server struct {
	DB     pgxpool.Pool
	Router *chi.Mux
}

func NewServer(Config conf.Config) Server {

	server := Server{
		DB:     createDb(Config.Datasource),
		Router: chi.NewRouter(),
	}

	return server
}

func createDb(Datasource conf.Datasource) pgxpool.Pool {

	conn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", Datasource.Username, Datasource.Password, Datasource.Host, Datasource.Port, Datasource.Database)
	pgxpool, err := pgxpool.Connect(context.Background(), conn)

	if err != nil {
		log.Fatal().Msg("cannot connect to the database")
	}
	return *pgxpool
}
