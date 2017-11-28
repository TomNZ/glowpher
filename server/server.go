package server

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

// NewHandler instantiates a new handler for serving the Glowpher frontend and API.
func NewHandler() http.Handler {
	mux := goji.NewMux()
	mux.HandleFunc(pat.New("/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	return mux
}
