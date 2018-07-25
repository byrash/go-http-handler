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
		fmt.Println(ioutil.ReadAll(request.Body))
		jsonDecoder := json.NewDecoder(request.Body)
		var paylod Payload
		err := jsonDecoder.Decode(&paylod)
		if err != nil {
			panic(err)
		}
		log.Printf("Paylod received %v", paylod)
		fmt.Fprintf(respWritter, "Hello %v", "Shivaji!!!")
	})
	http.ListenAndServe(":80", nil)
}
