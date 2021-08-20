package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func genericFetch(slug string) *http.Response {
	today := time.Now().Unix() // Default date
	fixedApiURL := fmt.Sprintf("%s/%s?date=%d", API_URL, slug, today)
	apiResponse, errResponse := http.Get(fixedApiURL)
	errorHandler(errResponse)
	return apiResponse
}
