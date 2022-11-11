package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	// "time"
	// "csv"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))

	for i, line := range lines {
		problems[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	return problems
}

func main() {
	correct := 0

	filePath := flag.String("csv", "problems.csv", "a csv file in the format o 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filePath)
	if err != nil {
		exit(fmt.Sprintf("There was an error while opening %s", *filePath))
	}

	r := csv.NewReader(file)
	records, readerErr := r.ReadAll()
	if readerErr != nil {
		exit(fmt.Sprintf("Theres was an error while reading %s", *filePath))
	}

	problems := parseLines(records)

	for _, problem := range problems {
		fmt.Printf("Question: %s?: ", problem.question)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == problem.answer {
			correct++
		}
	}

	fmt.Printf("Game over. Your score is %d out of %d\n", correct, len(problems))

}
