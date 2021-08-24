package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateServer() *Server {

	startDb(schemaObject)

	return &Server{
		Port:   os.Getenv("API_PORT"),
		Router: CreateRouter(),
	}

}

func (s *Server) Listen() {

	// Primera carga de hoy por default:
	todayDate := time.Now().Unix()
	dateData := fetchDayData(todayDate, queryProducts, queryOrigins)
	addDayData(dateData)

	log.Println("Default data uploaded -> Server ready to start")

	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Mount("/api", s.Router)
	mainRouter.Get("/", HandleRoot)

	serverError := http.ListenAndServe(s.Port, mainRouter)
	panicErrorHandler(serverError)

}
