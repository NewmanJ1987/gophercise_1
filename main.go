package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Get the name of the file to open
	fileName := flag.String("filename", "problems.csv", "The name of the csv file to read.")
	flag.Parse()
	// Open and read the contents of the file
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	// Put the csvs in a problems splice
	problems := parseProblems(records)
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for _, problem := range problems {
		fmt.Printf("%s\n", problem.question)
		input, _ := reader.ReadString('\n')
		if problem.answer == strings.TrimSuffix(input, "\n") {
			score++
		}
	}
	fmt.Printf("Your score is %d\n.", score)
}

type Problems struct {
	question string
	answer   string
}

func parseProblems(records [][]string) []Problems {
	problems := make([]Problems, len(records))
	for index, record := range records {
		problems[index].question = record[0]
		problems[index].answer = record[1]
	}
	return problems
}
