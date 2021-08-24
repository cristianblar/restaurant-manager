package server

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/cristianblar/restaurant-manager/api/lib"
	"github.com/cristianblar/restaurant-manager/api/utils"
	"github.com/go-chi/chi/v5"
)

type Pagination struct {
	TotalResults int         `dgraph:"Buyers.total,omitempty"`
	TotalPages   int         `dgraph:"Pages.total,omitempty"`
	PreviousPage string      `dgraph:"Page.previous,omitempty"`
	NextPage     string      `dgraph:"Page.next,omitempty"`
	Results      []lib.Buyer `dgraph:"results,omitempty"`
}

var currentDate string = time.Now().String()[0:10]

func HandleRoot(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to Restaurant Manager v1.0.0!"))
}

func HandleLoadData(res http.ResponseWriter, req *http.Request) {

	requestedDateString := req.URL.Query().Get("date")

	var response []byte

	if requestedDateString == "" {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		response = []byte(`{ "message": "Date in Unix timestamp required" }`)
	} else {
		requestedDate, convError := strconv.ParseInt(requestedDateString, 10, 64)
		if convError != nil {
			log.Println(convError.Error())
			http.Error(res, "Invalid Unix timestamp", http.StatusBadRequest)
			return
		}
		if time.Unix(requestedDate, 0).String()[0:10] == currentDate {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusBadRequest)
			response = []byte(`{ "result": "Date already synced" }`)
		} else {
			cleanError := lib.PrepareNewDate()
			if cleanError != nil {
				log.Println(cleanError.Error())
				http.Error(res, httpErrorMessage, http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			currentDate = time.Unix(requestedDate, 0).String()[0:10]
			newData := lib.FetchDayData(requestedDate, queryProducts, queryOrigins)
			lib.AddDayData(newData)
			res.WriteHeader(http.StatusOK)
			response = []byte(`{ "result": "Data synced" }`)
		}
	}

	res.Write(response)

}

func HandleAllBuyers(res http.ResponseWriter, req *http.Request) {

	requestedPageString := req.URL.Query().Get("page")

	var requestedPage int

	if requestedPageString == "" {
		requestedPage = 1
	} else {
		requestedPageConv, convError := strconv.ParseUint(requestedPageString, 10, 64)
		if convError != nil || requestedPageConv == 0 {
			requestedPage = 1
		} else {
			requestedPage = int(requestedPageConv)
		}
	}

	_, queryResult, queryError := lib.GetAllBuyers(queryAllBuyers)
	if queryError != nil {
		log.Println(queryError.Error())
		http.Error(res, httpErrorMessage, http.StatusInternalServerError)
		return
	}

	if queryResult == nil {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "message": "The synced date doesn't have buyers" }`))
		return
	} else {
		var (
			sliceLeft    int
			sliceRight   int
			previousPage string
			nextPage     string
		)
		resultsLength := len(queryResult.Q)
		totalPages := int(math.Ceil(float64(resultsLength) / 100))

		if requestedPage >= totalPages {
			requestedPage = totalPages
			nextPage = ""
			sliceRight = resultsLength
		} else {
			nextPage = fmt.Sprintf("%s/api/buyers?page=%d", currentDomain, requestedPage+1)
			sliceRight = (requestedPage * 100)
		}

		if requestedPage == 1 {
			sliceLeft = 0
			previousPage = ""
		} else {
			previousPage = fmt.Sprintf("%s/api/buyers?page=%d", currentDomain, requestedPage-1)
			sliceLeft = (requestedPage * 100) - 100
		}

		pagination := &Pagination{
			TotalResults: resultsLength,
			TotalPages:   totalPages,
			PreviousPage: previousPage,
			NextPage:     nextPage,
			Results:      queryResult.Q[sliceLeft:sliceRight],
		}
		jsonToSend, marshallError := utils.JsoniterMarshall(pagination, "dgraph")
		if marshallError != nil {
			log.Println(marshallError.Error())
			http.Error(res, httpErrorMessage, http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(jsonToSend)
	}

}

func HandleBuyerId(res http.ResponseWriter, req *http.Request) {

	buyerId := chi.URLParam(req, "buyerId")

	vars := map[string]string{"$id": buyerId}
	jsonResult, _, queryError := lib.GetBuyerById(queryBuyerById, vars)
	if queryError != nil {
		log.Println(queryError.Error())
		http.Error(res, httpErrorMessage, http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	if jsonResult == nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "message": "The requested ID doesn't exist" }`))
	} else {
		res.WriteHeader(http.StatusOK)
		res.Write(jsonResult)
	}

}
