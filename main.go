package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"url_shortener/middleware"
	"url_shortener/routes"
)

func run(
	stdout io.Writer,
) error {
	handler := http.NewServeMux()
	addr := ":8090"

	// middleware
	logger := log.New(stdout, "", log.LstdFlags)
	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	// routes
	if err := routes.RegisterRoutes(handler, logger); err != nil {
		return fmt.Errorf("RegisterRoutes: %w", err)
	}

	// start server
	server := http.Server{
		Addr:    addr,
		Handler: middlewareStack(handler),
	}

	go func() {
		log.Printf("listening on address %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("error run: ListenAndServe: %s\n", err)
		}
		log.Printf("stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 1)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("err run: Shutdown: %s\n", err)
	}
	log.Println("graceful shutdown complete")

	return nil
}

func main() {
	if err := run(
		os.Stdout,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
