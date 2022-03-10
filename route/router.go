package route

import (
	"encoding/json"
	"log"
	"main/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetListRoute(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")

	json.NewEncoder(response).Encode(service.Getlist())
}

func CreateRoute() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/book", GetListRoute).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
