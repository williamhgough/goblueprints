package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const BHT_URL = "http://words.bighugelabs.com/api/2/"

type BigHugh struct {
	APIKey string
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (b *BigHugh) Synonyms(term string) ([]string, error) {
	var syns []string
	url := fmt.Sprintf("%s/%s/%s/json", BHT_URL, b.APIKey, term)
	res, err := http.Get(url)
	if err != nil {
		return syns, errors.New("Bighugh: failed when looking for synonyms for\"" + term + "\" " + err.Error())
	}

	var data synonyms
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return syns, err
	}

	syns = append(syns, data.Noun.Syn...)
	syns = append(syns, data.Verb.Syn...)

	return syns, nil
}
