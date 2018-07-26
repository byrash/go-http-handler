package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Job represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var jobQueue chan Job

func main() {
	var maxJobQueueBuffer = 5
	jobQueue = make(chan Job, maxJobQueueBuffer)
	http.HandleFunc("/", handleRequest)
	dispatcher := NewDispatcher(MaxWorkers) // Creates worker pool
	dispatcher.Run(jobQueue)                // Watches for Job queue
	http.ListenAndServe(":8080", nil)
}

func handleRequest(respWritter http.ResponseWriter, request *http.Request) {

	// Only accept posts
	if !strings.EqualFold(request.Method, "POST") {
		respWritter.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		badRequestHandler(respWritter, err)
		return
	}
	log.Println(string(body))
	var paylodCollection PayloadCollection
	err = json.Unmarshal(body, &paylodCollection)
	if err != nil {
		badRequestHandler(respWritter, err)
		return
	}
	log.Printf("Paylod received %v", paylodCollection)

	for _, payload := range paylodCollection.Payloads {
		jobQueue <- Job{Payload: payload} // Put  job onto Job queue
	}
	respWritter.WriteHeader(http.StatusOK) // All good!!
}

func badRequestHandler(respWritter http.ResponseWriter, err error) {
	respWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respWritter.WriteHeader(http.StatusBadRequest)
	log.Printf("%v", err)
}
