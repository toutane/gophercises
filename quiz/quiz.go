package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "problems.csv", "File containing questions and answers.")
	timeLimit := flag.Int("l", 30, "The time limit in seconde.")
	flag.Parse()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		os.Exit(1)
	}

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Failes to read file: %v\n", err)
		os.Exit(1)
	}

	var correct int
	q, a := parseLines(data)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

problemsLoop:
	for i := 0; i < len(q); i++ {
		var answerCh = make(chan string)
		fmt.Printf("Problem #%d: %v = ", (i + 1), q[i])
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break problemsLoop
		case answer := <-answerCh:
			if answer == a[i] {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d corrects out of %d.\n", correct, len(q))
}

// ParseLines parse 2D data array into two arrays q for questions and a for answers.
func parseLines(data [][]string) ([]string, []string) {
	var q, a []string
	for _, arr := range data {
		q = append(q, arr[0])
		a = append(a, arr[1])
	}
	return q, a
}
