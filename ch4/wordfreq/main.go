// wordfreq counts the frequency of words in an input text file (histogram)
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//!+
func main() {

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	hist := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords) // break into words
	for input.Scan() {
		line := input.Text()
		hist[line]++ // zero-val
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
	for k, v := range hist {
		fmt.Printf("%s: %d\n", k, v)
	}
}

//!-
