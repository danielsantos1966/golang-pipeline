package main

import (
	"io"
	"log"
	"net/http"
    "time"
    "fmt"
    "os"
    "net"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
    t := time.Now().Format(time.RFC1123)
    u := uptime()
    h, err := os.Hostname()
    if err != nil {
        //panic(err)
        h = "unknown"
    }
    addrs, err := net.LookupHost(h)
    if err != nil {
        log.Printf("Cant get ipaddrr")
    }

    fmt.Fprintf(w, "Time: %s | host: %s | ipaddr: %s | uptime: %s | version: %s\n", t , h, addrs, u, version)
}

const version string = "2.2.2-222"

// VersionHandler handles incoming requests to /version
// and just returns a simple version number
func versionHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, version)
}

func main() {
	log.Printf("Listening on port 8080...")
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/status", statusHandler)
	http.ListenAndServe(":8080", nil)
}
