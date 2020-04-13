package main

import (
	"io"
	"log"
	"net/http"
)

const version string = "2.2.2-2"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func main() {
	log.Printf("Listening on port 8080...")
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8080", nil)
}
