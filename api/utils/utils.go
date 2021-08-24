package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"

	jsoniter "github.com/json-iterator/go"
)

func PanicErrorHandler(err error) {

	if err != nil {
		log.Panic(err.Error())
	}

}

func GenericFetch(slug string, date int64) *http.Response {

	fixedApiURL := fmt.Sprintf("%s/%s?date=%d", os.Getenv("API_URL"), slug, date)
	response, httpError := http.Get(fixedApiURL)
	PanicErrorHandler(httpError)

	return response

}

func JsoniterMarshall(v interface{}, tagKey string) ([]byte, error) {

	marshaller := jsoniter.Config{TagKey: tagKey}.Froze()
	return marshaller.Marshal(v)

}

func JsoniterUnmarshall(data []byte, v interface{}, tagKey string) error {

	unmarshaller := jsoniter.Config{TagKey: tagKey}.Froze()
	return unmarshaller.Unmarshal(data, v)

}
