package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/phyber/negroni-gzip/gzip"
	"github.com/subosito/gotenv"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"
	"gopkg.in/mgo.v2"
)

func init() {
	gotenv.Load()
}

func main() {
	var (
		addr  = flag.String("addr", ":8080", "Endpoint address")
		mongo = flag.String("mongo", os.Getenv("MONGO_URL"), "Mongo DB address, defaults to 'MONGO_URL' env var")
	)
	flag.Parse()

	log.Println("Dialing mongo...", *mongo)
	db, err := mgo.Dial(*mongo)
	if err != nil {
		log.Fatalln("Failed to connect to mongo", err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/polls", withCORS(withVars(withData(db, withAPIKey(handlePolls)))))

	n := negroni.Classic()
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(mux)

	log.Println("Starting web server on", *addr)
	graceful.Run(*addr, 2*time.Second, n)
	log.Println("Stopping...")
}

func withAPIKey(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isValidAPIKey(r.URL.Query().Get("key")) {
			respondErr(w, r, http.StatusUnauthorized, "invalid API key")
			return
		}
		fn(w, r)
	}
}

func isValidAPIKey(key string) bool {
	return key == "abc123"
}

func withData(d *mgo.Session, f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		thisDb := d.Copy()
		defer thisDb.Close()
		SetVar(r, "db", thisDb.DB("socialpoll"))
		f(w, r)
	}
}

func withVars(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		OpenVars(r)
		defer CloseVars(r)
		fn(w, r)
	}
}

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Location")
		fn(w, r)
	}
}
