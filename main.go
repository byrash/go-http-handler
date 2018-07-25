package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(respWritter http.ResponseWriter, request *http.Request) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
		var paylod Payload
		err = json.Unmarshal(body, &paylod)
		if err != nil {
			panic(err)
		}
		log.Printf("Paylod received %v", paylod)
		fmt.Fprintf(respWritter, "Hello %v", "Shivaji!!!")
	})
	http.ListenAndServe(":80", nil)
}
