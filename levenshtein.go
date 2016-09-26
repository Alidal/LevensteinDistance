package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// type Word struct {
// 	text string
// 	distance int
// }

func minOfThree(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		} else {
			return c
		}
	} else {
		if b < c {
			return b
		} else {
			return c
		}
	}
}

func levenshteinDistance(s1, s2 string) int {
	var lastdiag, olddiag int
	s1len := len(s1)
	s2len := len(s2)
	var column = make([]int, s1len+1)

	for i := 1; i < s1len; i++ {
		column[i] = i
	}
	for i := 1; i < s2len; i++ {
		column[0] = i
		lastdiag = i - 1
		for j := 1; j <= s1len; j++ {
			olddiag = column[j]
			diff := 1
			if s1[j-1] == s2[i-1] {
				diff = 0
			}
			column[j] = minOfThree(column[j]+1, column[j-1]+1, lastdiag+diff)
			lastdiag = olddiag
		}
	}
	return column[s1len]
}

func main() {
	// Open file
	file, err := os.Open("test_data.txt")
	// Handle file oper error
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create words array
	// var words []Word
	startWord := "test"

	// Read file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(levenshteinDistance(startWord, scanner.Text()))
	}
	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
