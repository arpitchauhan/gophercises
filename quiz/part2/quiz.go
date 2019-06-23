package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
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

func askQuestions(scanner *bufio.Scanner, questionsWithAnswers [][]string, score *int, done chan<- bool) {
	for index, questionsWithAnswer := range questionsWithAnswers {
		question, correctAnswer := questionsWithAnswer[0], questionsWithAnswer[1]

		fmt.Printf("Problem #%v: %v = ", index+1, question)
		scanner.Scan()

		userAnswer := scanner.Text()

		if userAnswer == correctAnswer {
			*score++
		}
	}

	done <- true
}

func conductQuiz(scanner *bufio.Scanner, questionsWithAnswers [][]string, timeLimit time.Duration) int {
	score := 0

	quizDone := make(chan bool)

	timer := time.NewTimer(timeLimit)

	go askQuestions(scanner, questionsWithAnswers, &score, quizDone)

	select {
	case <-quizDone:
		return score
	case <-timer.C:
		return score
	}
}

func main() {
	timeLimitPtr := flag.Int("limit", 30, "the time limit for quiz in seconds")
	csvFilePathPtr := flag.String("csv", "../problems.csv", "the path of the CSV file")

	flag.Parse()

	reader, err := getReaderFromFile(*csvFilePathPtr)
	check(err)

	records, err := extractRecordsFromReader(reader)
	check(err)

	userScore := conductQuiz(bufio.NewScanner(os.Stdin), records, time.Duration(*timeLimitPtr)*(time.Second))

	fmt.Printf("\nYou scored %v out of %v.", userScore, len(records))
}
