package main

import (
	"flag"
	"fmt"
	linkparser "github.com/gophercises/htmllinkparse"
	"os"
)

func main() {
	filename := flag.String("f", "ex1.html", "file to parse")

	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	links, err := linkparser.ParseHTML(file)
	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Printf("link: %s\ntext: %s\n", link.Href, link.Text)
	}
}
