package main

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
	"gopkg.in/mgo.v2"
)

func init() {
	gotenv.Load()
}

func main() {}

var db *mgo.Session

func dialdb() error {
	var err error
	log.Println("Dialing mongodb")
	db, err = mgo.Dial(os.Getenv("MONGO_URL"))
	return err
}

func closedb() {
	db.Close()
	log.Println("closed db connection")
}

type poll struct {
	Options []string
}

func loadOptions() ([]string, error) {
	var options []string
	iter := db.DB("socialpoll").C("polls").Find(nil).Iter()
	var p poll
	for iter.Next(&p) {
		options = append(options, p.Options...)
	}
	iter.Close()
	return options, iter.Err()
}
