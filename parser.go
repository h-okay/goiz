package main

import (
	"strings"
)

type ParseOptions struct {
	length int // amount of line to be parsed
}

func Parse(data [][]string, options ParseOptions) []Question {
	if options.length > 0 && options.length <= len(data) {
		return parseN(data, options.length)
	}
	return parseAll(data)
}

func parseN(data [][]string, n int) []Question {
	questions := make([]Question, 0, n)
	for i := 0; i < n; i++ {
		questions = append(questions, Question{
			id:       i + 1,
			question: data[i][0],
			answer:   strings.TrimSpace(data[i][1]),
		})
	}
	return questions
}

func parseAll(data [][]string) []Question {
	questions := make([]Question, 0, len(data))
	for i := 0; i < len(data); i++ {
		questions = append(questions, Question{
			id:       i + 1,
			question: data[i][0],
			answer:   strings.TrimSpace(data[i][1]),
		})
	}
	return questions
}
