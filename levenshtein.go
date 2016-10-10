package main

import (
	"bufio"
	"log"
	//"os"
	"sort"
	"github.com/pkg/profile"
	//"net/http"
	//"time"
	//"fmt"
	"io"
	//"time"
	//"fmt"
	"os"
	//"github.com/wblakecaldwell/profiler"
)

func levenshteinDistance(s1, s2 string) int {
	/***
		A little bit optimized Levenshtein algorithm, so it uses O(min(m,n))
		space instead of O(mn), where m and n - lengths of compared strings.
		The key observation is that we only need to access the contents
		of the previous column when filling the matrix column-by-column.
		Hence, we can re-use a single column over and over, overwriting its contents as we proceed.
	***/
	var prevDiagonalValue, buffer int
	s1len := len(s1)
	s2len := len(s2)

	// Initialize column
	var curColumn = make([]int, s1len+1)
	for i := 0; i < s1len; i++ {
		curColumn[i] = i
	}

	// Fill matrix column by column
	for i := 1; i <= s2len; i++ {
		curColumn[0] = i
		prevDiagonalValue = i - 1
		for j := 1; j <= s1len; j++ {
			// Set operation cost (all operations except match(M) has value 1)
			operationCost := 1
			if s1[j-1] == s2[i-1] {
				operationCost = 0
			}
			buffer = curColumn[j]
			curColumn[j] = minOfThree(curColumn[j]+1, curColumn[j-1]+1, prevDiagonalValue+operationCost)
			prevDiagonalValue = buffer
		}
	}
	return curColumn[s1len]
}

func run(startWord string, file io.Reader) Words{

	// Create words array
	var curWord Word
	var words Words

	// Read file word by word
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		curWord.Text = scanner.Text()
		curWord.Distance = levenshteinDistance(startWord, curWord.Text)
		words = append(words, curWord)
	}
	// Handle scanner errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Sort(words)

	return (words)
}

func createTestFile(possibleQuantity []int) {
	for _, value := range possibleQuantity {
		generateWords(value)
	}
}

func main() {
	// Works as a profiler
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()

	file, err := os.Open("test40000000.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//start := time.Now()

	run("test",file)

	//elapsed := time.Since(start)
	//fmt.Println("Task took ", elapsed)
}
