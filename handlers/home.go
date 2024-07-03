package handlers

import (
	"net/http"

	"github.com/windyboy/goth-starter/views/home" // Assuming this is the correct import path for 'home'
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, home.Index())
}
