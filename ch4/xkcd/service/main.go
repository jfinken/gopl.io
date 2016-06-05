package main

import (
	"fmt"
	"log"

	"github.com/jfinken/gopl.io/ch4/xkcd/index"
)

func main() {
	// open or create a new index
	index, err := index.NewIndex("xkcd.index")
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	//searchTerm := "bobby tables"
	searchTerm := "turing"
	// search!
	result, err := index.Search(searchTerm)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	//
	for i := 0; i < len(result.Hits); i++ {
		hit := result.Hits[i]
		fmt.Printf("ID: %s\tScore: %f\tFields: %v\n", hit.ID, hit.Score, hit.Fields)
	}
}
