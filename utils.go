package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func randomWord(n int) string {
	// Got code from here: http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

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

// Generate file filled with n random words.
func generateTestFileWithLength(n int) {

	fileName := fmt.Sprintf("test%v.txt", n)
	file, err := os.Create(fileName)
	// Handle file open error
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var randowWordLength int
	for i := 0; i < n; i++ {
		randowWordLength = rand.Intn(15)
		file.WriteString(randomWord(randowWordLength+1) + "\n")
		if math.Mod(float64(i), 1000000) == 0 && i != 0 {
			fmt.Println(i, "words has been created")
			// To check that program don't stuck
		}
	}
	fmt.Println(n, "words has been created")
}
