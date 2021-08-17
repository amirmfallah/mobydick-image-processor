package muxlib

import (
	"fmt"
	"net/http"
)

func initLog(r *http.Request) {
	fmt.Printf("%s-%s-%v-", r.Method, r.URL.Path, r.URL.Query())
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	HttpError404(w, "Requested resource doesn't exist. Please check your path.")
}

func ImageProcessHandler(w http.ResponseWriter, r *http.Request) {

}
