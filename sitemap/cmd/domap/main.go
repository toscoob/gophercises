package main

import (
	"flag"
	"fmt"
	"github.com/gophercises/sitemap"
	"log"
	"net/url"
)

func main() {
	rawUrl := flag.String("url", "", "url to examine")
	depth := flag.Uint("d", 0, "max link depth. Default unlimited")
	flag.Parse()

	_ = depth

	u, err := url.ParseRequestURI(*rawUrl)
	if err != nil {
		log.Fatal("Please provide valid url: ", err)
	}

	fmt.Println(u.Host)

	visited := make(map[string]struct{})
	err = sitemap.InspectURL(*u, visited, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	//TODO output to xml format
	fmt.Println("Collected links:")
	for l, _ := range visited{
		fmt.Println(l)
	}
}
