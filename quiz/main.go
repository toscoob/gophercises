package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	csvfile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	questions := 0
	answers := 0

	reader := bufio.NewReader(os.Stdin)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		questions++
		//fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		fmt.Printf("Question #%d: %s = ", questions, record[0])
		text, _ := reader.ReadString('\n')
		//fmt.Println(text)
		if strings.TrimRight(text, "\n") == record[1] {
			fmt.Println("Correct")
			answers++
		} else {
			fmt.Printf("Wrong. Correct answer: %s\n", record[1])
		}
	}

	csvfile.Close()

	fmt.Printf("Correct answers: %d/%d\n", answers, questions)
}
