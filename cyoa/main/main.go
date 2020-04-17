package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/gophercises/cyoa"
	"log"
	"net/http"
	"strconv"
	"html/template"
)

// go run main/main.go

func playStoryConsole(story map[string]cyoa.StoryEntry) {
	if arc, ok := story["intro"]; ok {
		for {
			fmt.Println(&arc)
			fmt.Println("Choose option:")
			var answer string
			_, err := fmt.Scanf("%s\n", &answer)
			if err != nil {
				panic(err)
			}
			if answer == "q" {
				fmt.Println("exiting...")
				return
			}
			//fmt.Printf("answer: %s\n", answer)
			madeChoice := false
			for i, opt := range arc.Options {
				//fmt.Printf("option: %s\n", string(i))
				if answer == strconv.Itoa(i+1) {
					//fmt.Printf("FOUND ANSWER: %s\n", string(i+1))
					arc = story[opt.Arc]
					madeChoice = true
					break
				}
			}
			if madeChoice == false {
				fmt.Println("Please choose one of options")
			}
			//break
		}
	} else {
		fmt.Println("Story intro not found")
	}
}



func main() {
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
	/*
		for k, v := range story {
			fmt.Printf("arc: %s\ncontents: %s\n", k, &v)
		}

	*/
	//playStoryConsole(story)
	tmpl := template.Must(template.ParseFiles("layout.html"))
	sh := &cyoa.StoryHandler{Story: story, CurrentArc: "intro", Tmpl: *tmpl}
	fmt.Println("Starting the server on :8080")
	http.Handle("/cyoa", sh)
	log.Fatal(http.ListenAndServe(":8080", nil))
	//_ = http.ListenAndServe(":8080", sh)
}
