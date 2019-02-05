package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/souvikhaldar/replayd/pkg/configloader"
)

var conf configloader.Config

func init() {
	log.Println("---replayd running---")
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	configloader.Load("/etc/replayd/config.json", &conf)
}

var b bytes.Buffer
var backup []string

// StoreInBuffer stores the incoming data to in-memory buffer
func StoreInBuffer(w http.ResponseWriter, r *http.Request) {
	log.Println("--Running in StoreInBuffer--")
	body, er := ioutil.ReadAll(r.Body)
	if er != nil {
		http.Error(w, er.Error(), 500)
		return
	}
	log.Println("Body recieved: ", string(body))
	backup = append(backup, string(body)+"\n")
	mux := &sync.Mutex{}
	mux.Lock()
	if _, e := b.Write(body); e != nil {
		http.Error(w, e.Error(), 500)
		return
	}
	b.WriteString("\n")
	mux.Unlock()
	fmt.Println("Inserted: ", string(body))
	fmt.Fprintf(w, "Successfully inserted data to buffer\n")
	return
}

// ReadFromBuffer responds with the in-memory buffer data
func ReadFromBuffer(w http.ResponseWriter, r *http.Request) {
	log.Println("--Running in ReadFromBuffer--")
	fmt.Fprint(w, b.String())
	return
}

func DeleteBuffer(w http.ResponseWriter, r *http.Request) {
	log.Println("--Running in DeleteBuffer--")
	buf := make([]byte, b.Len())
	if _, e := b.Read(buf); e != nil {
		http.Error(w, e.Error(), 500)
	}
	fmt.Fprint(w, fmt.Sprintf("Deleted: %s", string(buf)))
}

func Backup(w http.ResponseWriter, r *http.Request) {
	log.Println("--Reading backup data--")
	fmt.Fprintln(w, fmt.Sprintf("Rescued data: %s", backup))
	return
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", StoreInBuffer).Methods("POST")
	router.HandleFunc("/", ReadFromBuffer).Methods("GET")
	router.HandleFunc("/", DeleteBuffer).Methods("DELETE")
	router.HandleFunc("/backup", Backup).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+conf.Port, router))
}
