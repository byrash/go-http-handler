package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
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
		var paylod PayloadCollection
		err = json.Unmarshal(body, &paylod)
		if err != nil {
			badRequestHandler(respWritter, err)
			return
		}
		log.Printf("Paylod received %v", paylod)
		fmt.Fprintf(respWritter, "Hello %v", "Shivaji!!!")
	})
	http.ListenAndServe(":80", nil)
}

func badRequestHandler(respWritter http.ResponseWriter, err error) {
	respWritter.Header().Set("Content-Type", "application/json; charset=UTF-8")
	respWritter.WriteHeader(http.StatusBadRequest)
	log.Printf("%v", err)
}
