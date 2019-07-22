package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/MichaelKaaden/redux-server-golang/counters"
	"github.com/gin-gonic/gin"
)

type CounterResponse struct {
	Counter counters.Counter `json:"counter"`
}

type CounterData struct {
	Data    CounterResponse `json:"data"`
	Message string          `json:"message"`
	Status  int             `json:"status"`
}

type CountersResponse struct {
	Counters counters.Counters `json:"counters"`
}

type CountersData struct {
	Data    CountersResponse `json:"data"`
	Message string           `json:"message"`
	Status  int              `json:"status"`
}

type IncDecBody struct {
	By int `json:"by"`
}

var theCounters = counters.New()

// GetCounters returns all counters.
func GetCounters(context *gin.Context) {
	data := buildCountersResponse(theCounters)
	context.JSON(http.StatusOK, data)
}

// GetCounter returns one specific counter. If it doesn't yet
// exist, it creates it.
func GetCounter(context *gin.Context) {
	index := getIdFromUrl(context)

	c := counters.GetCounter(theCounters, index)
	data := buildCounterResponse(c)
	context.JSON(http.StatusOK, data)
}

// PutCounter sets a counter to a new value.
func PutCounter(context *gin.Context) {
	index := getIdFromUrl(context)

	newValue, err := getIntValueFromBody(context, "count")
	if err != nil {
		log.Fatal(err)
	}

	c, err := counters.SetCounter(theCounters, index, newValue)
	if err != nil {
		log.Fatal(err)
	}
	data := buildCounterResponse(c)
	context.JSON(http.StatusOK, data)
}

// IncrementCounter increments the counter specified in the URL
// with the increment specified in the body.
func IncrementCounter(context *gin.Context) {
	index := getIdFromUrl(context)

	by, err := getByFromBody(context)
	if err != nil {
		log.Fatal(err)
	}

	c, err := counters.Increment(theCounters, index, by)
	if err != nil {
		log.Fatal(err)
	}
	data := buildCounterResponse(c)
	context.JSON(http.StatusOK, data)
}

// DecrementCounter decrements the counter specified in the URL
// with the decrement specified in the body.
func DecrementCounter(context *gin.Context) {
	index := getIdFromUrl(context)

	by, err := getByFromBody(context)
	if err != nil {
		log.Fatal(err)
	}

	c, err := counters.Decrement(theCounters, index, by)
	if err != nil {
		log.Fatal(err)
	}
	data := buildCounterResponse(c)
	context.JSON(http.StatusOK, data)
}

func buildCounterResponse(cntr counters.Counter) CounterData {
	response := CounterResponse{Counter: cntr}
	return CounterData{Data: response, Status: http.StatusOK, Message: "okay"}
}

func buildCountersResponse(cntrs *counters.Counters) CountersData {
	response := CountersResponse{Counters: *cntrs}
	return CountersData{Data: response, Status: http.StatusOK, Message: "okay"}
}

// getIntValueFromBody reads a value from the request's body and returns it as an int value.
func getIntValueFromBody(context *gin.Context, key string) (int, error) {
	request := context.Request
	err := request.ParseForm()
	if err != nil {
		return 0, err
	}
	value := request.FormValue(key)
	return strconv.Atoi(value)
}

// getByFromBody reads the "by" key of a JSON sent in the request's body.
func getByFromBody(context *gin.Context) (int, error) {
	b, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	// make sure the body is going to be closed
	defer func() {
		err := context.Request.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var idb IncDecBody
	err = json.Unmarshal(b, &idb)
	if err != nil {
		// log.Fatal(err)

		fmt.Println("error:", err)
		return 0, nil
	}

	return idb.By, nil
}

// getIdFromUrl returns the ID specified in the URL.
func getIdFromUrl(context *gin.Context) int {
	id := context.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Couldn't convert %#v to int\n", id)
	}
	return index
}

func printAsJSON(data IncDecBody) {
	var jsonData []byte
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("jsonData: %v\n", string(jsonData))
	fmt.Printf("jsonData: %#v\n", string(jsonData))
}
