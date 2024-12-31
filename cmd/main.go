package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Mount("/books", BookRoutes())
	http.ListenAndServe(":3000", r)
}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	bh := Bookhandler{}
	r.Get("/", bh.GetAllBooks)
	r.Get("/{id}", bh.GetOneBook)
	r.Post("/", bh.CreateBook)
	r.Delete("/{id}", bh.DeleteBook)
}
