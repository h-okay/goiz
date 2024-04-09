package main

import (
	"fmt"
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

	for _, question := range quiz.questions {
		timer := time.NewTimer(time.Duration(quiz.timer.secondsPerQuestion) * time.Second)
		fmt.Printf("Problem #%d: %s = ", question.id, question.question)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("Time limit exceed!")
			continue
		case answer := <-answerCh:
			if answer == question.answer {
				quiz.score++
				fmt.Println("Correct!")
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", quiz.score, len(quiz.questions))
}
