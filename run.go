package main

import (
	"fmt"
	//"google.golang.org/appengine"
	//"google.golang.org/appengine/log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleRedirection)
	http.ListenAndServe(":8080", nil)
}

// Handle redirection to somewhere
func handleRedirection(w http.ResponseWriter, r *http.Request) {
	//ctx := appengine.NewContext(r)

	links := GetLinks()
	name := strings.TrimLeft(r.URL.Path, "/")[1:]
	url, ok := links[name]

	if !ok {
		http.Error(w, fmt.Sprintf("Not found / %s", name), http.StatusNotFound)
		return
	}


	fmt.Printf("Redirect to %s\n", url)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
