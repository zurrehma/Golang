package main

import (
	"flag"
	"fmt"
	"os"
	// "reflect"
	"encoding/csv"
	"time"
	"sync"
)

type problem struct {
	question, answer string
}
var wg sync.WaitGroup
func main()  {
	defer finshExam()
	csvFilename := flag.String("csv","problems.csv","A csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "time for quiz in seconds")
	flag.Parse()
	// fmt.Println(time.Duration(*timeLimit) * time.Second)
	file, err := os.Open(*csvFilename)
	if err != nil {
		error(fmt.Sprintf("Failed to open file: %s\n", *csvFilename))
	}
	
	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		error(fmt.Sprintf("Failed to parse file%s/n", *csvFilename))
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	takeQuiz(problems, timer)
}

func finshExam() {
	fmt.Println("Exam Finished")
}

func takeQuiz(problems []problem, timer *time.Timer) {
	correct := 0	
	for index, problem := range problems {
		select {
		case <-timer.C:
			fmt.Println("Over")
			fmt.Printf("You scored %d out of %d\n", correct, len(problems))
			return
		default:
		fmt.Printf("Problem #%d: %s =", index+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct ++
		}
	}
	}
	fmt.Printf("You scored %d out of %d\n", correct, len(problems))	
}

func parseLines(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for index, line := range lines {
		problems[index] = problem{
			question : line[0],
			answer: line[1],
		}
	}
	return problems
}

func error(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}