package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var b bytes.Buffer

// StoreInBuffer stores the incoming data to in-memory buffer
func StoreInBuffer(w http.ResponseWriter, r *http.Request) {
	log.Println("--Running in StoreInBuffer--")
	body, er := ioutil.ReadAll(r.Body)
	if er != nil {
		http.Error(w, er.Error(), 500)
		return
	}
	log.Println("Body recieved: ", string(body))
	mux := &sync.Mutex{}
	mux.Lock()
	if _, e := b.Write(body); e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	mux.Unlock()
	fmt.Println("Inserted: ", string(body))
	fmt.Fprintf(w, "Successfully inserted data to buffer")
	return
}

// ReadFromBuffer responds with the in-memory buffer data
func ReadFromBuffer(w http.ResponseWriter, r *http.Request) {
	log.Println("--Running in ReadFromBuffer--")
	fmt.Fprintf(w, b.String())
	return
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", StoreInBuffer).Methods("POST")
	router.HandleFunc("/", ReadFromBuffer).Methods("GET")
	log.Fatal(http.ListenAndServe(":8192", router))
}
