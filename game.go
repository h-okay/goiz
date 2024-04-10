package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
	"time"
)

type Timer struct {
	secondsPerQuestion int
}

type Question struct {
	id       int
	question string
	answer   string
}

type Quiz struct {
	score     uint8
	questions []Question
	timer     Timer
}

func Play(quiz Quiz) {
	fmt.Println("You have", quiz.timer.secondsPerQuestion, "seconds per question.")
	input := make(chan string, len(quiz.questions))

	for _, question := range quiz.questions {
		ctx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(quiz.timer.secondsPerQuestion)*time.Second,
		)
		defer cancel()

		fmt.Printf("Problem #%d: %s = ", question.id, question.question)
		go getInput(input)

		select {
		case <-ctx.Done():
			fmt.Println("Time limit exceeded!")
		case answer := <-input:
			if answer == question.answer {
				quiz.score++
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", quiz.score, len(quiz.questions))
}

func getInput(input chan string) {
	in := bufio.NewReader(os.Stdin)
	for {
		result, err := in.ReadString('\n')
		if err != nil {
			return
		}
		input <- strings.TrimSpace(result)
	}
}
