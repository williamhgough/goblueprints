package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln("can't create server")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Worldy friendy"))
}
