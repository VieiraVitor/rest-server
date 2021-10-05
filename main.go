package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/vieiravitor/rest-server/handler"
)

const addr = ":8080"

func main() {
	handler := handler.Handler{}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Route("/transaction/v1", func(router chi.Router) {
		router.Post("/send", handler.SaveTransaction)
	})
	fmt.Printf("Server started at port %s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal("could not initialize the server")
	}
}
