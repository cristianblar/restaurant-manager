package main

import (
	"fmt"
	"log"
	"net/http"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func genericFetch(slug string, date int64) *http.Response {
	fixedApiURL := fmt.Sprintf("%s/%s?date=%d", API_URL, slug, date)
	apiResponse, errResponse := http.Get(fixedApiURL)
	errorHandler(errResponse)
	return apiResponse
}
