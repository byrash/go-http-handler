package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var Queue chan Payload

func main() {
	var MAX_QUEUE_BUFFER = 5
	Queue = make(chan Payload, MAX_QUEUE_BUFFER)
	http.HandleFunc("/", func(respWritter http.ResponseWriter, request *http.Request) {

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
			Queue <- payload
		}
		respWritter.WriteHeader(http.StatusOK) // All good!!
	})
	go startUpload()
	http.ListenAndServe(":80", nil)

}

func startUpload() {
	for { // For ever
		select { // Synchronously: Wait for paylods to be added to queue to process
		case paylod := <-Queue:
			err := paylod.Upload()
			if err != nil {
				log.Printf("Error uploading Paylod %v", err)
			}
		}
	}
}

func badRequestHandler(respWritter http.ResponseWriter, err error) {
	respWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respWritter.WriteHeader(http.StatusBadRequest)
	log.Printf("%v", err)
}
