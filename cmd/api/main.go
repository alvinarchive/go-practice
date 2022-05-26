package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"movie.alvintanoto.id/internal/data"

	_ "github.com/lib/pq"
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
	db   struct {
		dsn string
	}
}

// Define an application struct to hold depedencies for HTTP handlers, helpers and middleware
type application struct {
	config config
	logger *log.Logger
	models data.Models
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
	flag.StringVar(&cfg.db.dsn, "db-dsn", "postgres://postgres:postgres@localhost/moviedb?sslmode=disable", "Postgresql DSN")
	flag.Parse()

	// initialize logger
	// prefixed with date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// init db
	db, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	// db connected
	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	// migration in code
	// although this works, tightly coupling the execution of migrations with your application source code can be problematic in the long term

	// migrationDriver, err := postgres.WithInstance(db, &postgres.Config{})
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// migrator, err := migrate.NewWithDatabaseInstance("file:///Users/alvintanoto/Documents/sandbox/gosandbox/movie/migrations", "postgres", migrationDriver)
	// if err != nil {
	// 	logger.Fatal(err)
	// }

	// err = migrator.Up()
	// if err != nil && err != migrate.ErrNoChange {
	// 	fmt.Printf("err sana")
	// 	logger.Fatal(err)
	// }

	// logger.Printf("database migration applied")

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
	err = srv.ListenAndServe()
	logger.Fatal(err)

	fmt.Println("Hello World!")
}

// database
func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	// create context with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// use pingcontext to establish connection to db
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
