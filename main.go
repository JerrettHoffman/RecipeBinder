package main

import (
	"RecipeBinder/auth"
	db "RecipeBinder/internal/db/dbtest"
	"RecipeBinder/router"

	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	}

	userArgs := ""
	if len(os.Args) > 1 {
		userArgs = strings.ToLower(os.Args[1])
	}

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

		auth.Setup()

		// Initialize the server
		goServer := &http.Server{
			Addr:                         ":8080",
			Handler:                      auth.SessionMiddleware(r.Handler),
			DisableGeneralOptionsHandler: false,
			ReadTimeout:                  10 * time.Second,
			ReadHeaderTimeout:            10 * time.Second,
			WriteTimeout:                 10 * time.Second,
			IdleTimeout:                  0,
			MaxHeaderBytes:               1 << 20,
			ErrorLog:                     &log.Logger{},
		}

		// Setup listener for interrupt signal (run in goroutine)
		idleConnsClosed := make(chan struct{})
		go func() {
			sigint := make(chan os.Signal, 1)
			signal.Notify(sigint, os.Interrupt)
			<-sigint

			// We received an interrupt signal, shut down.
			if err := goServer.Shutdown(context.Background()); err != nil {
				// Error from closing listeners, or context timeout:
				log.Printf("HTTP server Shutdown: %v", err)
			}
			close(idleConnsClosed)
		}()

		log.Print("Starting server...\n\n\nhttp://localhost:8080/search\n\n")

		// Kickoff server
		if err := goServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe failed: %v", err)
		}

		<-idleConnsClosed
	}

}
