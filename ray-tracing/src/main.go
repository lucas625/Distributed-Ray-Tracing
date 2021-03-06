package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/lucas625/Distributed-Ray-Tracing/ray-tracing/src/rest"
	"log"
	"net/http"
	"time"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/path-tracing", rest.RunPathTracing)

	server := &http.Server{
		Handler:      router,
		Addr:         ":8081",
		WriteTimeout: 1 * time.Hour,
		ReadTimeout:  1 * time.Hour,
	}

	fmt.Println("Server running!")

	err := server.ListenAndServe()
	log.Fatal(err)
}
