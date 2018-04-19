package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/williamhgough/goblueprints/thesaurus"
)

func init() {
}

func main() {
	thesaurus := &thesaurus.BigHugh{APIKey: "788ebfc31d7ce57fec47c8932b52f60d"}

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
