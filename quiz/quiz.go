package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("problems.csv")
	check(err)

	r := csv.NewReader(file)

	records, err := r.ReadAll()
	check(err)

	scanner := bufio.NewScanner(os.Stdin)

	score := 0

	for index, line := range records {
		question, correctAnswer := line[0], line[1]

		fmt.Printf("Problem #%v: %v = ", index+1, question)
		scanner.Scan()

		userAnswer := scanner.Text()
		if userAnswer == correctAnswer {
			score++
		}
	}

	fmt.Printf("You scored %v out of %v.", score, len(records))
}
