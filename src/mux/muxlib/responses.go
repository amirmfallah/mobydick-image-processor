package muxlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	lib "image-processor-dokkaan.ir/lib"
)

func HttpOptionsResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
}

func jsonResponseHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")
}

func failSafeError(w http.ResponseWriter) {
	w.Write([]byte(`{
			"@type":"Error",
			"code":500,
			"title":"Internal error",
			"description":"Internal server error."
			}`))
}

func HttpSuccessResponse(w http.ResponseWriter, statusCode int, payload []byte) {
	jsonResponseHeaders(w)
	w.WriteHeader(statusCode)
	if statusCode == http.StatusNoContent {
		return
	}
	_, err := w.Write(payload)
	if err != nil {
		fmt.Print("HttpSuccessResponse-Write Error:", err)
		failSafeError(w)
	}
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

func HttpError400(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusBadRequest, "Bad request", description)
}

func HttpError401(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusUnauthorized, "Unauthorized", description)
}

func HttpError402(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusPaymentRequired, "Payment Required", description)
}

func HttpError403(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusForbidden, "Forbidden", description)
}

func HttpError404(w http.ResponseWriter, description string) {
	HttpErrorResponse(w, http.StatusNotFound, "Not found", description)
}

func HttpError500(w http.ResponseWriter) {
	HttpErrorResponse(w, http.StatusInternalServerError, "Internal error", "Internal server error.")
}

func HttpError503(w http.ResponseWriter) {
	HttpErrorResponse(w, http.StatusServiceUnavailable, "Service Unavailable", "Service is unavailable.")
}

func HttpErrorWith(w http.ResponseWriter, err error) {
	if err == nil {
		fmt.Print("HttpErrorWith: Error was nil.")
		failSafeError(w)
		return
	}
	if errors.Is(err, lib.ErrNotFound) {
		HttpError404(w, err.Error())
	} else if errors.Is(err, lib.ErrForbidden) {
		HttpError403(w, err.Error())
	} else if errors.Is(err, lib.ErrDatabaseConnection) {
		HttpError500(w)
	} else {
		HttpError500(w)
	}
}

func HttpHandleTokenErrors(w http.ResponseWriter, err error) {
	if err == nil {
		fmt.Print("HttpHandleTokenErrors: Error was nil.")
		failSafeError(w)
		return
	}
	message := err.Error()
	if message != "Token is expired" {
		message = "Unauthorized token detected. Please sign-out and sign-in again." +
			" If that didn't help, please contact your administrator."
	}
	HttpError401(w, message)
}
