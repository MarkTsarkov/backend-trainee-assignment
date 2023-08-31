package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/marktsarkov/avito/internal/config"
	"github.com/marktsarkov/avito/internal/storage"
	"github.com/marktsarkov/avito/internal/transport"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func run() error {
	// read config from env
	cfg := config.Read()

	// start dbpool
	dbpool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// create segment repository
	segmentStorage := storage.NewStorage(dbpool)

	// Create Table
	segmentStorage.CreateTables(context.Background())

	// create http server with application injected
	httpServer := transport.NewHttpServer(segmentStorage)

	// create http router
	router := mux.NewRouter()
	router.HandleFunc("/segment", httpServer.CreateSegment).Methods("POST")
	router.HandleFunc("/segment/{slug}", httpServer.DeleteSegment).Methods("DELETE")
	router.HandleFunc("/user/{userid}", httpServer.AddAndRemoveSegmentsOnUser).Methods("POST", "DELETE")
	router.HandleFunc("/user/{userid}", httpServer.ShowUserSegments).Methods("GET")

	srv := &http.Server{
		Addr:    cfg.HTTPAddr,
		Handler: router,
	}

	// listen to OS signals and gracefully shutdown HTTP server
	stopped := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("HTTP Server Shutdown Error: %v", err)
		}
		close(stopped)
	}()

	log.Printf("Starting HTTP server on %s", cfg.HTTPAddr)

	// start HTTP server
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe Error: %v", err)
	}

	<-stopped

	log.Printf("Have a nice day!")

	return nil
}
