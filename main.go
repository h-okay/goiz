package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csv := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	secondsPerQuestion := flag.Int("time", 3, "time limit per question, in seconds")
	amountOfQuestions := flag.Int("amount", 5, "amount of questions to be loaded from the csv file, if invalid all file contents are loaded.")
	flag.Parse()

	data, err := LoadFile(LoadOptions{filepath: *csv})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	timer := Timer{secondsPerQuestion: *secondsPerQuestion}
	questions := Parse(data, ParseOptions{length: *amountOfQuestions})
	quiz := Quiz{questions: questions, timer: timer}

	Play(quiz)
}
