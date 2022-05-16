package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Declare version as global variable
const version = "1.0.0"

// Define a config struct to hold configuration setting for our app
// port -> network port
// env -> string (development, staging, production)
// configuration setting from command line when app starts
type config struct {
	port int
	env  string
}

// Define an application struct to hold depedencies for HTTP handlers, helpers and middleware
type application struct {
	config config
	logger *log.Logger
}

type envelope map[string]interface{}

func main() {
	// Declare config instance
	var cfg config

	// Read the value of command line flag
	// default port: 3000
	// default env: development
	flag.IntVar(&cfg.port, "port", 3000, "API Server Port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// initialize logger
	// prefixed with date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare a new servemux and add a /v1/healthcheck route to the healthCheckHandler method
	// comment this because we're using httprouter
	// mux := http.NewServeMux()
	// mux.HandleFunc("/v1/healthcheck", app.healthCheckHandler)

	// Declare HTTP server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("Starting %s server on %d", cfg.env, cfg.port)
	err := srv.ListenAndServe()
	logger.Fatal(err)

	fmt.Println("Hello World!")
}
