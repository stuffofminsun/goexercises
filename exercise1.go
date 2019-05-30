package main

// Gophercise 1: Create a program that will read in a quiz provided via a CSV file (more details below) and will then
// give the quiz to a user keeping track of how many questions they get right and how many they get incorrect.
// Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

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
	// Reads the questions
	problems := parseQuiz("problems.csv")

	// start the quiz
	correct := 0
	for _, p := range problems {
		fmt.Printf("Problem: %s?\n", p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correct++
			fmt.Printf("Correct! You have %d answer(s) right so far\n", correct)
		} else {
			fmt.Printf("Wrong! You have %d answers(s) right so far\n", correct)
		}
	}

	fmt.Printf("You scored %d out of %d problems correctly", correct, len(problems))
}

func parseQuiz(filename string) []problem {
	file, _ := os.Open(filename)
	fr := csv.NewReader(bufio.NewReader(file))

	var ret []problem
	for {
		line, error := fr.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		q := strings.TrimSpace(line[0])
		a := strings.TrimSpace(line[1])
		build := problem{question: q, answer: a}
		ret = append(ret, build)
	}
	return ret
}

type problem struct {
	question string
	answer   string
}
