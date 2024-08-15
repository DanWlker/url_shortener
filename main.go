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

	"github.com/DanWlker/url_shortener/middleware"
	"github.com/DanWlker/url_shortener/routes"
	"github.com/DanWlker/url_shortener/storage"
	// "github.com/jackc/pgx/v5/pgxpool"
	// "github.com/redis/go-redis/v9"
)

func run(
	ctx context.Context,
	stdout io.Writer,
) error {
	handler := http.NewServeMux()
	addr := ":8090"

	// middleware
	logger := log.New(stdout, "", log.LstdFlags)
	middlewareStack := middleware.CreateStack(
		middleware.Logging,
	)

	// mock
	storageClient := storage.NewMockStorage()

	// redis
	// client := redis.NewClient(
	// 	&redis.Options{
	// 		Addr:     "localhost:6379",
	// 		Password: "",
	// 		DB:       0,
	// 	},
	// )
	// storageClient := storage.NewRedisClient(ctx, client)

	// postgres
	// db_url, ok := os.LookupEnv("DATABASE_URL")
	// if !ok {
	// 	return fmt.Errorf("os.LookupEnv: Cannot find DATABASE_URL in environment")
	// }
	//
	// db, err := pgxpool.New(ctx, db_url)
	// if err != nil {
	// 	return fmt.Errorf("pgx.Connect: %w", err)
	// }
	// defer db.Close() // TODO: I still don't know if db close should close before or after shutdown
	//
	// storageClient := storage.NewPostgresClient(ctx, db)

	if err := storageClient.Ping(); err != nil {
		return fmt.Errorf("storageClient.Ping: %w", err)
	}

	// routes
	if err := routes.RegisterRoutes(handler, logger, storageClient); err != nil {
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
	defer cancel() // TODO: Not sure if this should happen before or after db shutdown

	if err := server.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("err run: Shutdown: %s\n", err)
	}

	log.Println("graceful shutdown complete, pending defers")

	return nil
}

func main() {
	if err := run(
		context.Background(),
		os.Stdout,
	); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
