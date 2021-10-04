package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const addr = ":9090"

func main() {
	router := chi.NewRouter()
	router.Route("/transaction/v1", func(r chi.Router) {
		r.Post("/send", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome"))
		})
	})
	fmt.Printf("Server started at port %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("could not initialize the server")
	}
}
