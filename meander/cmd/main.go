package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/subosito/gotenv"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
	"github.com/williamhgough/goblueprints/meander"
)

func init() {
	gotenv.Load()
}

func main() {
	// Set max CPUS
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Allow for optional port
	var addr = flag.String("addr", ":8080", "Endpoint address")
	flag.Parse()
	// Set API key
	meander.APIKey = os.Getenv("PLACES_API")

	// Create new Mux
	mux := http.NewServeMux()

	// Serve different available journeys
	mux.HandleFunc("/journeys", withCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		respond(w, r, meander.Journeys)
	}))

	// Fetch available recommendations
	mux.HandleFunc("/recommendations", withCORS(func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		q.Lat, _ = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		q.Lng, _ = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		q.Radius, _ = strconv.Atoi(r.URL.Query().Get("radius"))
		q.CostRangeStr = r.URL.Query().Get("cost")
		places := q.Run()
		respond(w, r, places)
	}))

	// Use negroni for gzip and logging.
	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)

	// Boostrap server
	log.Println("Starting web server on", *addr)
	graceful.Run(*addr, 2*time.Second, n)
	log.Println("Stopping...")
}

// respond is generic helper to send response with encoded JSON data.
func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

// Wrap a handler with CORS support
func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}
