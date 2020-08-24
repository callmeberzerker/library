package server

import "net/http"

func (srv *Server) InitAuthorRoutes() {

	r := srv.Router

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

}
