package main

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/a-h/templ"
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

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/about", http.StatusPermanentRedirect)
	})

	router.Handle("/about", templ.Handler(templates.About(portfolio.EducationList)))

	languageList := make([]string, 0)
	for _, item := range portfolio.JobList {
		languageList = append(languageList, item.Languages...)
	}

	router.Get("/experience", func(w http.ResponseWriter, r *http.Request) {
		skills := r.URL.Query().Get("skill")
		templates.Experience(getXpListFiltered(skills), removeDuplicate[string](languageList), skills).Render(r.Context(), w)
	})

	router.Handle("/projects", templ.Handler(templates.Projects(portfolio.ProjectList)))

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func getXpListFiltered(skill string) []*portfolio.JobExperienceItemProps {
	if skill == "" {
		return portfolio.JobList
	}

	l := make([]*portfolio.JobExperienceItemProps, 0)
	for _, v := range portfolio.JobList {
		if slices.ContainsFunc[[]string, string](v.Languages, func(s string) bool {
			return strings.ToLower(s) == strings.ToLower(skill)
		}) {
			l = append(l, v)
		}
	}

	return l
}

func removeDuplicate[T comparable](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
