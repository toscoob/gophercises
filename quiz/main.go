package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

//go build . && ./quiz

func main() {
	rand.Seed(time.Now().UnixNano())

	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	shuffle := flag.Bool("shuffle", false, "shuffle problems")
	duration := flag.Int("d", 30, "quiz duration")
	flag.Parse()

	csvfile, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	if *shuffle {
		fmt.Println("Shuffle problems")
		shuffleProblems(problems)
	}

	correct := 0

	fmt.Println("Press enter to start test")
	fmt.Scanf("\n")
	timer := time.NewTimer(time.Duration(*duration) * time.Second)

problemloop:
	for i, p := range problems {
		fmt.Printf("Question #%d: %s =\n", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("Correct")
				correct++
			} else {
				fmt.Println("Wrong")
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func shuffleProblems(problems []problem) {
	rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
