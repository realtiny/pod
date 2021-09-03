package main

import (
	"fmt"
	"net/http"

	"github.com/realtiny/pod"
)

// Handler turn http.Handler to pod.Handler
type Handler struct {
	http.Handler
}

// NewHandler new Handler
func NewHandler(h http.Handler) *Handler {
	return &Handler{h}
}

// ServeHTTP implement pod.Handler
func (h *Handler) ServeHTTP(
	w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	h.Handler.ServeHTTP(w, r)
	next(w, r)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "Hello")
	})

	app := pod.New()
	app.Push(NewHandler(mux))

	app.Run(":8000")
}
