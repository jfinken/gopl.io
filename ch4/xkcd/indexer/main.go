package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/jfinken/gopl.io/ch4/xkcd/index"
)

const (
	xkcdURL       = "https://xkcd.com/"
	xkcdPostfix   = "/info.0.json"
	numGoroutines = 80
	maxXkcd       = 1692 // 2016-06-10
)

// Comic contains the data for an XKCD comic minus the PNGs
type Comic struct {
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
}

func indexComic(index *index.Index, in chan int, syncGroup *sync.WaitGroup) {

	defer syncGroup.Done()

	// Note: for loop on a channel will break on channel close
	// curl https://xkcd.com/327/info.0.json
	for num := range in {
		url := xkcdURL + strconv.Itoa(num) + xkcdPostfix

		fmt.Printf("Indexing: %s\n", url)

		time.Sleep(10 * time.Millisecond)

		resp, err := http.Get(url)

		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		// presume nil err obviates check here?
		if resp != nil {
			defer resp.Body.Close()
		}

		// unmarshal data
		var comic Comic
		if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
		// index it
		err = index.Add(strconv.Itoa(comic.Num), comic)
		if err != nil {
			fmt.Printf("%s\n", err.Error())
			continue
		}
	}
}

func main() {

	// open or create a new index
	index, err := index.NewIndex("xkcd.index")
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}

	in := make(chan int)

	// Start producing over the numeric xkcd JSON urls: 1 - 1689
	go func() {
		for i := 1; i <= maxXkcd; i++ {
			// send the xkcd num to the chan
			in <- i
		}
		close(in)
	}()

	// Setup and run indexing consumers
	syncGroup := new(sync.WaitGroup)
	syncGroup.Add(numGoroutines)
	for j := 0; j <= numGoroutines; j++ {
		go indexComic(index, in, syncGroup)
	}
	syncGroup.Wait()
}
