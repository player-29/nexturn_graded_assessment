// Exercise 4: Online Examination System 
// Topics Covered: Go Loops, Go Break and Continue, Go Constants, Go Strings, Go 
// Functions, Go Errors 
// Case Study: 
// Develop an online examination system where users can take a quiz. 
// 1. Question Bank: Define a slice of structs to store questions. Each question 
// should have a question string, options (array), and the correct answer. 
// 2. Take Quiz: Use loops to iterate over questions and display them one by one. 
// Allow the user to select an answer by entering the option number. 
// o Use continue to skip invalid inputs and prompt the user again. 
// o Use break to exit the quiz early if the user enters a specific command 
// (e.g., "exit"). 
// 3. Score Calculation: After the quiz, calculate the user's score and display it. Use 
// conditions to classify performance (e.g., "Excellent", "Good", "Needs 
// Improvement"). 
// 4. Error Handling: Handle errors like invalid input during the quiz (e.g., entering a 
// non-integer value for an option). 
// Bonus: 
// • Add a timer for the quiz, limiting each question to a fixed amount of time.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Question struct {
	Question string
	Options  []string
	Answer   int // The correct option index (1-based)
}

var questionBank = []Question{
	{
		Question: "What is the capital of France?",
		Options:  []string{"1. Berlin", "2. Madrid", "3. Paris", "4. Rome"},
		Answer:   3,
	},
	{
		Question: "Which programming language is used for Go?",
		Options:  []string{"1. Java", "2. Python", "3. Go", "4. C++"},
		Answer:   3,
	},
	{
		Question: "What is 2 + 2?",
		Options:  []string{"1. 3", "2. 4", "3. 5", "4. 6"},
		Answer:   2,
	},
}

func main() {
	fmt.Println("Welcome to the Online Examination System")
	fmt.Println("Type 'exit' to quit the quiz at any time.")

	score := takeQuiz()
	classifyPerformance(score, len(questionBank))
}

func takeQuiz() int {
	score := 0
	reader := bufio.NewReader(os.Stdin)

	for i, question := range questionBank {
		fmt.Printf("\nQuestion %d: %s\n", i+1, question.Question)
		for _, option := range question.Options {
			fmt.Println(option)
		}

		answerCh := make(chan int, 1)
		errorCh := make(chan error, 1)

		go func() {
			for {
				fmt.Print("Enter your choice (1-4): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)

				if strings.ToLower(input) == "exit" {
					close(answerCh)
					close(errorCh)
					return
				}

				choice, err := strconv.Atoi(input)
				if err != nil || choice < 1 || choice > 4 {
					fmt.Println("Invalid input. Please enter a number between 1 and 4.")
					continue
				}

				answerCh <- choice
				close(answerCh)
				close(errorCh)
				return
			}
		}()

		select {
		case answer := <-answerCh:
			if answer == question.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Wrong!")
			}
		case <-time.After(10 * time.Second): // Timer of 10 seconds per question
			fmt.Println("\nTime's up! Moving to the next question.")
			continue
		}
	}

	return score
}

func classifyPerformance(score, total int) {
	fmt.Printf("\nYour score: %d/%d\n", score, total)
	percentage := float64(score) / float64(total) * 100

	switch {
	case percentage >= 80:
		fmt.Println("Performance: Excellent!")
	case percentage >= 50:
		fmt.Println("Performance: Good")
	default:
		fmt.Println("Performance: Needs Improvement")
	}
}
