package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// StoreInBuffer stores the incoming data to in-memory buffer
func StoreInBuffer(w http.ResponseWriter, r *http.Request) {

}

// ReadFromBuffer responds with the in-memory buffer data
func ReadFromBuffer(w http.ResponseWriter, r *http.Request) {

}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", StoreInBuffer).Methods("POST")
	router.HandleFunc("/", ReadFromBuffer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8192", router))
}
