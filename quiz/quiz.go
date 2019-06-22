package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func getReaderFromFile(filename string) (io.ReadCloser, error) {
	reader, err := os.Open(filename)
	return reader, err
}

func extractRecordsFromReader(reader io.ReadCloser) ([][]string, error) {
	defer reader.Close()
	return csv.NewReader(reader).ReadAll()
}

func conductQuiz(scanner *bufio.Scanner, questionsWithAnswers [][]string) int {
	score := 0

	for index, questionsWithAnswer := range questionsWithAnswers {
		question, correctAnswer := questionsWithAnswer[0], questionsWithAnswer[1]

		fmt.Printf("Problem #%v: %v = ", index+1, question)
		scanner.Scan()

		userAnswer := scanner.Text()

		if userAnswer == correctAnswer {
			score++
		}
	}

	return score
}

func main() {
	reader, err := getReaderFromFile("problems.csv")
	check(err)

	records, err := extractRecordsFromReader(reader)
	check(err)

	userScore := conductQuiz(bufio.NewScanner(os.Stdin), records)

	fmt.Printf("\nYou scored %v out of %v.", userScore, len(records))
}
