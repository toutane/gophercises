package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	var filename string
	flag.StringVar(&filename, "f", "problems.csv", "File containing questions and answers.")
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
	q, a := parse(data)

	for i := 0; i < len(q); i++ {
		var input string
		fmt.Printf("Problem #%d: %v = ", (i + 1), q[i])
		fmt.Scan(&input)
		if input == a[i] {
			correct++
		}
	}
	fmt.Printf("You scored %d corrects out of %d.\n", correct, len(q))
}

// Parse parse 2D data array into two arrays q for questions and a for answers.
func parse(data [][]string) ([]string, []string) {
	var q, a []string
	for _, arr := range data {
		q = append(q, arr[0])
		a = append(a, arr[1])
	}
	return q, a
}
