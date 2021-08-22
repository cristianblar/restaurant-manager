package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jsoniter "github.com/json-iterator/go"
)

func errorHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}

func genericFetch(slug string, date int64) *http.Response {
	API_URL := os.Getenv("API_URL")
	fixedApiURL := fmt.Sprintf("%s/%s?date=%d", API_URL, slug, date)
	apiResponse, errResponse := http.Get(fixedApiURL)
	errorHandler(errResponse)
	return apiResponse
}

func jsoniterMarshall(v interface{}, tagKey string) []byte {
	marshaller := jsoniter.Config{TagKey: tagKey}.Froze()
	bytes, marshallError := marshaller.Marshal(v)
	errorHandler(marshallError)
	return bytes
}

func jsoniterUnmarshall(data []byte, v interface{}, tagKey string) {
	unmarshaller := jsoniter.Config{TagKey: tagKey}.Froze()
	unmarshallError := unmarshaller.Unmarshal(data, v)
	errorHandler(unmarshallError)
}
