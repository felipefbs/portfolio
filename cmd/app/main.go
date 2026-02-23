package main

import (
	"log"
	"net/http"

	"github.com/felipefbs/portfolio/portfolio"
	"github.com/felipefbs/portfolio/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	fileHandler := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileHandler))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		templates.Home(portfolio.EducationList, portfolio.JobList, portfolio.ProjectList).Render(r.Context(), w)
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
