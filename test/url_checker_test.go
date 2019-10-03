package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/lookup/redis"
	"github.com/url-checker/handler"
)

var err error

// Assign a given URL to test whether it is a malicious website on basis of the database
func TestUrlUnsafeFromDb(t *testing.T) {
	conn := redis.GetPool().Get()
	_, err = conn.Do("HMSET", "www.example.com", "url", "www.example.com", "status", "Unsafe")
	if err != nil {
		log.Fatal(err)
	}

	url := string("www.example.com")

	req, err := http.NewRequest("GET", "/"+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `{"URL":"www.example.com","Status":"Unsafe"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}

// Assign a given URL to test whether it is a safe website on basis of the database
func TestUrlSafeFromDb(t *testing.T) {
	conn := redis.GetPool().Get()
	_, err = conn.Do("HMSET", "www.example1.com", "url", "www.example1.com", "status", "Safe")
	if err != nil {
		log.Fatal(err)
	}

	url := string("www.example1.com")

	req, err := http.NewRequest("GET", "/"+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `{"URL":"www.example1.com","Status":"Safe"}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}

// Assign a given URL to test whether it is not in the database, considered to be unknown
func TestUrlNotInDb(t *testing.T) {
	url := string("www.example2.com")

	req, err := http.NewRequest("GET", "/"+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `null` + "\n" + `Unknown url: not found in DB`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}

// Do not assign any URL to test whether this app gives an error message to let user provide URL
func TestGivenUrlEmpty(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `null` + "\n" + `No url is given, please provide URL`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}
