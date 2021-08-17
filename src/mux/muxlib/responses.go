package muxlib

import (
	"encoding/json"
	"fmt"
	"net/http"

	lib "image-processor-dokkaan.ir/lib"
)

func jsonResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
}

func HttpError404(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusNotFound, "Not found", description)
}

func failSafeError(w http.ResponseWriter) {
	w.Write([]byte(`{
			"@type":"Error",
			"code":500,
			"title":"Internal error",
			"description":"Internal server error."
			}`))
}

func HttpErrorResponse(w http.ResponseWriter, statusCode int, title, description string) {
	jsonResponseHeaders(w)
	w.WriteHeader(statusCode)
	errorResponse := lib.NewErrorResponse(statusCode, title, description)
	err := json.NewEncoder(w).Encode(errorResponse)
	if err != nil {
		fmt.Print("HttpErrorResponse-json.Encode Error:", err)
		failSafeError(w)
	}
}
