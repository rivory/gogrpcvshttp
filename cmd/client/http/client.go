package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type helloWorld struct {
	Message string `json:"message"`
}

func main() {
	hw := helloWorld{Message: "tototest"}
	postBody, _ := json.Marshal(&hw)
	responseBody := bytes.NewBuffer(postBody)
	start := time.Now()
	resp, err := http.Post("http://localhost:8080/hello", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("http.Post took %v microseconds \n", elapsed)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)
}

// TODO: big request body

// TODO: embricated body
