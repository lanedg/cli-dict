package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Definition struct {
	Word     string `json:"word"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
		}
	} `json:"meanings"`
}

func main() {
	wordIn := "yes"
	res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + wordIn)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("dict api error")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var defs []Definition

	err = json.Unmarshal(body, &defs)
	if err != nil {
		panic(err)
	}

	def := defs[0]
	meanings := def.Meanings
	fmt.Print(def.Word + "\n\n")

	for _, meaning := range meanings {
		fmt.Printf("%s: \n\n", meaning.PartOfSpeech)
		for _, defin := range meaning.Definitions {
			fmt.Printf("%s\n\n", defin.Definition)
			fmt.Printf("Example : %s\n\n", defin.Example)
		}
		fmt.Printf("___________________________\n\n")
	}
}
