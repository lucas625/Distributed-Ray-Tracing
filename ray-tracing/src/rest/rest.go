package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// runPathTracing runs the requested path tracing, sending a matrix of colors as response.
//
// Parameters:
// 	responseWriter - The response writer.
// 	request        - The request.
//
// Returns:
// 	none
//
func runPathTracing(responseWriter http.ResponseWriter, request *http.Request) {
	var data map[string]interface{}
	bodyAsBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(responseWriter, "Failed to decode your request.", 500)
	}
	err = json.Unmarshal(bodyAsBytes, &data)
	if err != nil {
		http.Error(responseWriter, "Failed to parse your request.", 500)
	}

	fmt.Printf("%v", data)
	responseWriter.Write([]byte("ok!"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/path-tracing", runPathTracing)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 1 * time.Hour,
		ReadTimeout:  1 * time.Hour,
	}

	fmt.Println("Server running!")

	err := server.ListenAndServe()
	log.Fatal(err)
}
