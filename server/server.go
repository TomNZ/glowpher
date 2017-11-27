package server

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func NewHandler() http.Handler {
	mux := goji.NewMux()
	mux.HandleFunc(pat.New("/"), func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})
	return mux
}
