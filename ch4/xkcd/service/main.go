package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jfinken/gopl.io/ch4/xkcd/index"
)

// TODO: minor DRY
const xkcdURL = "https://xkcd.com/"

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("USAGE: %s search terms...\n", os.Args[0])
		os.Exit(1)
	}

	// open or create a new index
	index, err := index.NewIndex("xkcd.index")
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	searchTerm := strings.Join(os.Args[1:], " ")
	// search!
	result, err := index.Search(searchTerm)
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	//
	for i := 0; i < len(result.Hits); i++ {
		hit := result.Hits[i]
		//fmt.Printf("ID: %s\tScore: %f\tFields: %v\n", hit.ID, hit.Score, hit.Fields)
		fmt.Printf("%s%s\n", xkcdURL, hit.ID)
	}
}
