package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadFromBuffer(t *testing.T) {
	req := httptest.NewRequest("POST", "/", bytes.NewBufferString("souvik"))
	res := httptest.NewRecorder()
	handler := http.HandlerFunc(StoreInBuffer)
	handler.ServeHTTP(res, req)
	if status := res.Code; status != 200 {
		t.Error("Response of Storing is not OK, the response recieved is: ", status)
	}
	RequiredBody := "Successfully inserted data to buffer"
	if gotString := res.Body.String(); gotString != RequiredBody {
		t.Error("Respose body of Storing does'nt match the required response. Recieved: ", gotString)
	}
	request := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	secHandler := http.HandlerFunc(ReadFromBuffer)
	secHandler.ServeHTTP(rr, request)
	if stat := rr.Code; stat != 200 {
		t.Error("Response code of reading is not OK, response code: ", stat)
	}
	if bod := rr.Body.String(); bod != "souvik" {
		t.Error("Response body not correct; Response body: ", bod)
	}
}
