package handlers

import (
	"net/http"

	"log/slog"

	"github.com/a-h/templ"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request) error

func Process(h HttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Processing request", "method", r.Method, "url", r.URL)
		if err := h(w, r); err != nil {
			slog.Error("Error processing request", "err", err)
		}
	}
}

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
}
