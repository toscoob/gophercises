package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/gophercises/cyoa"
)

// go run main/main.go

// write console player which shows story entry and waits for input

func main(){
	jsonFilename := flag.String("j", "scenario.json", "json file with scenario")

	flag.Parse()

	jsonContent, err := ioutil.ReadFile(*jsonFilename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.ReadStoryJSON(jsonContent)
	if err != nil {
		panic(err)
	}

	for k, v := range story {
		fmt.Printf("arc: %s\ntitle: %s\n", k, v.Title)
	}
}