package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"../model"
)

var validatingURL string

func validateUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// URL handler processes HTTP GET request
func UrlHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	validatingURL = "http://" + url

	if validateUrl(validatingURL) == true {
		validatedurl := strings.Split(r.URL.Path, "/")[1]
		lookup, err := model.GetURL(validatedurl)
		fmt.Println(url)
		// Write the url details as application/json to the client
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(lookup)

		if len(validatedurl) == 0 {
			fmt.Fprintf(w, "No url is given, please provide URL")
			return
		} else if err == model.NoUrlFound {
			fmt.Fprintf(w, "Unknown url: cannot found in db")
			return
		}
	} else if validateUrl(validatingURL) == false {
		fmt.Fprintf(w, "cannot validate url")
	}

}
