package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/subosito/gotenv"
	"github.com/williamhgough/goblueprints/thesaurus"
)

func init() {
	gotenv.Load()
}

func main() {
	apiKey := os.Getenv("BHT_API_KEY")
	thesaurus := &thesaurus.BigHugh{APIKey: apiKey}

	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		word := s.Text()
		syns, err := thesaurus.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for word: "+word, err)
		}

		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms")
		}

		for _, syn := range syns {
			fmt.Println(syn)
		}
	}
}
