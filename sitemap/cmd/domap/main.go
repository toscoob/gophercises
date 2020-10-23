package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/sitemap"
	"log"
	"net/url"
	"os"
)

//call example:
//go run cmd/domap/main.go --url https://calhoun.io -d 2

func main() {
	rawUrl := flag.String("url", "", "url to examine")
	depth := flag.Uint("d", 0, "max link depth. Default unlimited")
	flag.Parse()

	u, err := url.ParseRequestURI(*rawUrl)
	if err != nil {
		log.Fatal("Please provide valid url: ", err)
	}

	fmt.Println(u.Host)

	visited := make(map[string]struct{})
	err = sitemap.InspectURL(*u, visited, 1, *depth)
	if err != nil {
		log.Fatal(err)
	}

	byteXML, err := sitemap.ComposeXML(u.Host, visited)
	if err != nil {
		log.Fatal(err)
	}
	//todo write to file
	//fmt.Println(string(byteXML))
	f, err := os.Create("out.xml") // create/truncate the file
	if err != nil { log.Fatal(err) } // panic if error
	defer f.Close() // make sure it gets closed after
	fmt.Fprintln(f, string(byteXML))
	fmt.Println(string(byteXML))


	//fmt.Println("Collected links:")
	//for l, _ := range visited{
	//	fmt.Println(l)
	//}
}
