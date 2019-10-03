package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"../model"
)

// URL handler processes HTTP GET request
func UrlHandler(w http.ResponseWriter, r *http.Request) {

	url := strings.Split(r.URL.Path, "/")[1]
	lookup, err := model.GetURL(url)

	// Write the url details as application/json to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(lookup)

	if len(url) == 0 {
		fmt.Fprintf(w, "No url is given, please provide URL")
		return
	} else if err == model.NoUrlFound {
		fmt.Fprintf(w, "Unknown url: not found in DB")
		return
	}
}
