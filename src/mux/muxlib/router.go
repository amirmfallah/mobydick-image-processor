package muxlib

import (
	"net/http"

	"github.com/gorilla/mux"
)

func GetMuxRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/api/v1/images/product", ProductImageProcessHandler).
		Methods(http.MethodPut, http.MethodOptions)

	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	r.Use(mux.CORSMethodMiddleware(r))

	return r
}
