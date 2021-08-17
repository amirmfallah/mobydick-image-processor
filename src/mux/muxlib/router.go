package muxlib

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetMuxRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/sharfire/tags", nil).
		Methods(http.MethodGet, http.MethodOptions)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
