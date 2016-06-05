package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/jfinken/gopl.io/ch4/xkcd/index"
)

const (
	xkcdURL       = "https://xkcd.com/"
	xkcdPostfix   = "/info.0.json"
	numGoroutines = 80
	maxXkcd       = 1689
)

/*
curl https://xkcd.com/327/info.0.json
var data = `
{
    "month": "10",
    "num": 327,
    "link": "",
    "year": "2007",
    "news": "",
    "safe_title": "Exploits of a Mom",
    "transcript":
    "[[A woman is talking on the phone, holding a cup]]\nPhone: Hi, this is your son's school. We're having some computer trouble.\nMom: Oh dear\u00c3\u00a2\u00c2\u0080\u00c2\u0094did he break something?\nPhone: In a way\u00c3\u00a2\u00c2\u0080\u00c2\u0094\nPhone: Did you really name your son \"Robert'); DROP TABLE Students;--\" ?\nMom: Oh, yes. Little Bobby Tables, we call him.\nPhone: Well, we've lost this year's student records. I hope you're happy.\nMom: And I hope you've learned to sanitize your database inputs.\n{{title-text: Her daughter is named Help I'm trapped in a driver's license factory.}}",
    "alt": "Her daughter is named Help I'm trapped in a driver's license factory.",
    "img": "http:\/\/imgs.xkcd.com\/comics\/exploits_of_a_mom.png",
    "title": "Exploits of a Mom",
    "day": "10"
}
*/

// Comic contains the data for an XKCD comic
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
	for num := range in {
		url := xkcdURL + strconv.Itoa(num) + xkcdPostfix
		fmt.Printf("Indexing: %s\n", url)

		time.Sleep(500 * time.Millisecond)
		/*
			resp, err := http.Get(url)
			defer resp.Body.Close()

			if err != nil {
				fmt.Printf("%s\n", err.Error())
				continue
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
		*/
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
