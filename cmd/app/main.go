package main

import (
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/a-h/templ"
	"github.com/felipefbs/portfolio/portfolio"
	"github.com/felipefbs/portfolio/templates"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/reload", func(w http.ResponseWriter, r *http.Request) {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			err := websocket.Message.Send(ws, "reload")
			if err != nil {
				fmt.Println(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				fmt.Println(err)
			}
		}).ServeHTTP(w, r)
	})

	fileHandler := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fileHandler))

	router.Handle("/", templ.Handler(templates.About(portfolio.EducationList)))

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
