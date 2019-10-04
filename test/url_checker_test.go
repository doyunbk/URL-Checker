package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"../handler"
	"../model"
)

var err error

// Assign URL to test whether it is an invalidated url
func TestInvalidateUrl(t *testing.T) {
	url := string(";zc3b:-$`www.validateurl.com/zv/?bceq**&dvcse/")

	req, err := http.NewRequest("GET", "/"+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `Cannot validate url`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}

// Assign URL to test whether it is an unsafe website on basis of the database
func TestUrlUnsafeFromDb(t *testing.T) {
	conn := model.GetPool().Get()
	_, err = conn.Do("HMSET", "www.example.com", "url", "www.example.com", "status", "Unsafe")
	if err != nil {
		log.Fatal(err)
	}

	url := string("www.example.com/&qed?cxvvczd#&/z&32d")

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

// Assign URL to test whether it is a safe website on basis of the database
func TestUrlSafeFromDb(t *testing.T) {
	conn := model.GetPool().Get()
	_, err = conn.Do("HMSET", "www.example1.com", "url", "www.example1.com", "status", "Safe")
	if err != nil {
		log.Fatal(err)
	}

	url := string("www.example1.com/?cvde#?bcx34g1dwe/zcv~@#asz/")

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

// Assign URL to test whether it is not in the database, considered to be unknown url
func TestUrlNotInDb(t *testing.T) {
	url := string("www.example2.com/&fvcx$233/vcds!?54dza")

	req, err := http.NewRequest("GET", "/"+url, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `null` + "\n" + `Unknown url: cannot found in db`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}

// Do not assign any URL to test whether this app gives an error message to provide URL
func TestGivenUrlEmpty(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.UrlHandler)

	handler.ServeHTTP(rr, req)

	expected := `No url is given, please provide URL`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Handler returned unexpected values: received %v but expected %v",
			rr.Body.String(), expected)
	}
}
