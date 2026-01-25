package main

import (
	"RecipeBinder/router"
	"log"
	"net/http"
	"time"
)

func main() {
	// Initialize the server state
	r := router.Router{}
	r.Setup()

	// Initialize the server
	goServer := &http.Server{
		Addr:                         ":8080",
		Handler:                      r.Mux,
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		ReadHeaderTimeout:            10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  0,
		MaxHeaderBytes:               1 << 20,
		ErrorLog: &log.Logger{},
	}

	// Kickoff server
	if err := goServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe failed: %v", err)
	}
}
