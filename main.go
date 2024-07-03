package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/windyboy/goth-starter/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error("Error loading .env file", "err", err)
		log.Fatal(err)

	}
	router := chi.NewRouter()
	router.Handle("/*", public())
	router.Get("/", handlers.Process(handlers.HandleHome))
	// fmt.Println("Hello, world!")
	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("Starting server...", "listenAddr", listenAddr)
	if err := http.ListenAndServe(listenAddr, router); err != nil {
		slog.Error("Error starting server", "err", err)
		log.Fatal(err)
	}
}
