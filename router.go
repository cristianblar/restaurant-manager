package main

import "github.com/go-chi/chi/v5"

func CreateRouter() *chi.Mux {

	router := chi.NewRouter()

	router.Get("/load-data", HandleLoadData)
	router.Get("/buyers", HandleAllBuyers)
	router.Get("/buyers/{buyerId}", HandleBuyerId)

	return router

}
