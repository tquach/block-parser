package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tquach/block-parser/parser"
)

var (
	targetURL = flag.String("targetURL", "http://localhost:9000", "URl of server with data to parse")
)

func main() {
	flag.Parse()

	client := http.Client{}
	resp, err := client.Get(*targetURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	content := parser.Parse(body)
	log.Printf("Decoded content: %q\n", content)
}
