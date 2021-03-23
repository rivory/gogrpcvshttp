package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rivory/gogrpcvshttp/domain"
)

// Hhttp compose the http handler
type Hhttp struct {
	service domain.LogicInterface
}

// ProvideHTTPHandler provides http handler implementation
func ProvideHTTPHandler() Hhttp {
	return Hhttp{
		service: domain.LogicService{},
	}
}

type helloWorld struct {
	Message string `json:"message"`
}

// Handle handles
func (handler Hhttp) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// do nothing
	}
	d := json.NewDecoder(r.Body)
	hw := helloWorld{}
	err := d.Decode(&hw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	log.Printf("Http Received: %v", hw.Message)

	hw.Message = handler.service.Uppercase(hw.Message)
	res, _ := json.Marshal(hw)

	w.Write(res)
}
