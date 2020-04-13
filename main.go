package main

import (
	"io"
	"log"
	"net/http"
    "time"
    "fmt"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

const version string = "1.2.3-4"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
    t := time.Now().Format(time.RFC1123)
    u := uptime()
    //Body := "The current time is:"
fmt.Fprintf(w, "Time: %s | uptime: %s | version: ", t , u)


	io.WriteString(w, version)
}

func main() {
	log.Printf("Listening on port 8080...")
	http.HandleFunc("/version", versionHandler)
	http.ListenAndServe(":8080", nil)
}
