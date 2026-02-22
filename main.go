package main

import (
	db "RecipeBinder/internal/db/dbtest"
	"RecipeBinder/router"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	}

	userArgs := strings.ToLower(os.Args[1])

	if userArgs == "dbtest" {
		println("About to run dbtest")
		err := db.DbTest()
		if err != nil {
			println("DBTest failed with the following error: " + err.Error())
		}
	} else {
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
			ErrorLog:                     &log.Logger{},
		}

		// Kickoff server
		if err := goServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}
	}
}
