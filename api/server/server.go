package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cristianblar/restaurant-manager/api/lib"
	"github.com/cristianblar/restaurant-manager/api/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

type Server struct {
	Port   string
	Router *chi.Mux
}

func CreateServer() *Server {

	lib.StartDb(schemaObject)

	return &Server{
		Port:   os.Getenv("API_PORT"),
		Router: CreateRouter(),
	}

}

func (s *Server) Listen() {

	// Primera carga de hoy por default:
	todayDate := time.Now().Unix()
	dateData := lib.FetchDayData(todayDate, queryProducts, queryOrigins)
	lib.AddDayData(dateData)

	log.Println("Default data uploaded -> Server ready to start")

	mainRouter := chi.NewRouter()
	mainRouter.Use(middleware.Logger)
	mainRouter.Use(middleware.Recoverer)
	mainRouter.Use(cors.Default().Handler)
	mainRouter.Mount("/api", s.Router)
	mainRouter.Get("/", HandleRoot)

	serverError := http.ListenAndServe(s.Port, mainRouter)
	utils.PanicErrorHandler(serverError)

}
