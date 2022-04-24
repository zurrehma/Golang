package main

import (
	"flag"
	"fmt"
	"os"
	// "reflect"
	"encoding/csv"
	"time"
)

type problem struct {
	question, answer string
}

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
	answerCh := make(chan string)
// v2.1
forloop:
	for index, problem := range problems {
		fmt.Printf("Problem #%d: %s =", index+1, problem.question)	
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <-answer
		}()
		select {
		case <-timer.C:
			// v2
			// fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
			// we can use return to break the loop or we can also do it with break or continue using lables (https://www.ardanlabs.com/blog/2013/11/label-breaks-in-go.html)
			// (https://stackoverflow.com/questions/46792159/labels-break-vs-continue-vs-goto) 
			// v2
			// return 
			// v2.1
			break forloop
		case answer := <-answerCh:
			if answer == problem.answer {
				correct ++
			}
	}
	}
	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))	
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