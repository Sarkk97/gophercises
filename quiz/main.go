/**
- Read Quiz from CSV file
- Quiz reads each question to the user
- Quiz keeps track of correct and incorrect answers
- Regardless of correctness, next question should be answered immediately
- At end of quiz, output total number of correct answers

USING TIMER
- Press enter to start timer
- When timer ends, quiz should stop
*/
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type quizResult struct {
	questions int
	correct   int
	incorrect int
}

var result quizResult

var quizFile string
var quizDuration time.Duration
var quizShuffle bool

func init() {
	const (
		defaultQuiz    = "problems.csv"
		quizUsage      = "Path of quiz file (in csv)"
		durationUsage  = "Quiz duration. Value should be suffixed with appr unit e.g h,m,s,ms,us,ns"
		defaultShuffle = false
		shuffleUsage   = "Shuffle the Quiz questions"
	)
	defaultDuration, _ := time.ParseDuration("30s")
	flag.StringVar(&quizFile, "quizfile", defaultQuiz, quizUsage)
	flag.DurationVar(&quizDuration, "duration", defaultDuration, durationUsage)
	flag.BoolVar(&quizShuffle, "shuffle", defaultShuffle, shuffleUsage)
}

func displayResult() {
	fmt.Printf("You got %d correct out of %d questions\n", result.correct, result.questions)
}

func runQuiz(records [][]string, ch chan bool) {
	if quizShuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(records), func(i, j int) {
			records[i], records[j] = records[j], records[i]
		})
	}
	result.questions = len(records)
	for _, q := range records {
		fmt.Printf("%s= ", q[0])
		var ans string
		_, err := fmt.Scanln(&ans)
		if err != nil {
			log.Fatal(err)
		}
		if ans == q[1] {
			result.correct++
		} else {
			result.incorrect++
		}
	}
	ch <- true
}

func main() {
	//Read quiz from csv file (problems.txt)
	fmt.Println("Reading Quiz questions...")
	flag.Parse()
	done := make(chan bool)

	f, err := os.Open(quizFile)

	if err != nil {
		log.Fatalf("An error occured: %v", err)
	}

	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()

	fmt.Printf("Press Enter to start quiz. The duration is %s: (Enter)", quizDuration)
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		log.Fatal(err)
	}

	if char == '\n' {
		timer := time.NewTimer(quizDuration)
		fmt.Println("Your time starts now!")
		go func(ch chan<- bool) {
			<-timer.C
			fmt.Println("\nYour time is up!")
			ch <- true
		}(done)

		go runQuiz(records, done)
	}

	<-done
	displayResult()
}
