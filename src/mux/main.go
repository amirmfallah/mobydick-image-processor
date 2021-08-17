package main

import (
	"fmt"
	"net/http"

	muxlib "image-processor-dokkaan.ir/muxlib"
)

func init() {
	fmt.Printf("Mux cold start-")
}

func main() {
	StartServer()
}

func StartServer() {
	router := muxlib.GetMuxRouter()
	localPort := "0.0.0.0:3211"
	fmt.Println("Starting local server at", localPort, "...")
	http.ListenAndServe(localPort, router)
}
