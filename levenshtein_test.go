package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

type testpairMin struct {
	values []int
	min    int
}

type testpairRandom struct {
	length         int
	expectedLength int
}

type testpairLines struct {
	count         int
	expectedCount int
}

type testpairlevenshteinDistance struct {
	s1       string
	s2       string
	distance int
}

var testsMin = []testpairMin{
	{[]int{1, 2}, 1},
	{[]int{50, 100}, 50},
	{[]int{-1, 1}, -1},
}

var testsRand = []testpairRandom{
	{1, 1},
	{10, 10},
	{3, 3},
}

var testsGenerate = []testpairLines{
	{10, 10},
	{15, 15},
	{52, 52},
}

var words = []Words{
	{{"test", 4},
		{"testa", 5},
		{"testab", 6},
		{"testabc", 7},
		{"testabcd", 8},
		{"testabcde", 9}},
}

var testlevenshteinDistance = []testpairlevenshteinDistance{
	{"test", "test", 0},
	{"test", "testa", 1},
	{"test", "testab", 2},
	{"test", "testabc", 3},
	{"test", "testabcd", 4},
	{"test", "testabcde", 5},
}

var testRun = Words{
	{"test", 0},
	{"testa", 1},
	{"testab", 2},
	{"testabc", 3},
	{"testabcd", 4},
	{"testabcde", 5},
}

func TestMinOfThree(t *testing.T) {
	for _, pair := range testsMin {
		v := minOfThree(pair.values[0], pair.values[1], pair.min)
		if v != pair.min {
			t.Error(
				"For", pair.values,
				"expected", pair.min,
				"got", v,
			)
		}
	}
}

func TestRandomWord(t *testing.T) {
	for _, pair := range testsRand {
		v := randomWord(pair.length)
		if len(v) != pair.expectedLength {
			t.Error(
				"For", pair.length,
				"expected", pair.expectedLength,
				"got", v,
			)
		}
	}
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func TestGenerateTestFileWithLength(t *testing.T) {
	for _, pair := range testsGenerate {
		generateTestFileWithLength(pair.count)
		r, err := os.Open("test" + strconv.Itoa(pair.count) + ".txt")
		if err != nil {
			log.Fatal(err)
		}
		defer r.Close()
		v, _ := lineCounter(r)
		if v != pair.expectedCount {
			t.Error(
				"For", pair.count,
				"expected", pair.expectedCount,
				"got", v,
			)
		}
	}
}

func TestLength(t *testing.T) {
	for _, pair := range words {
		v := pair.Len()
		if v != len(pair) {
			t.Error(
				"For", len(pair),
				"expected", len(pair),
				"got", pair.Len(),
			)
		}
	}
}

func TestLess(t *testing.T) {
	for _, pair := range words {
		for i, _ := range pair {
			v := pair.Less(0, i)

			if v != (pair[0].Distance < pair[i].Distance) {
				t.Error(
					"For", pair[0].Distance < pair[i].Distance,
					"expected", pair[0].Distance < pair[i].Distance,
					"got", pair.Less(0, i),
				)
			}
		}
	}
}

func TestLevensteinDistance(t *testing.T) {
	for _, pair := range testlevenshteinDistance {
		v := levenshteinDistance(pair.s1, pair.s2)
		if v != pair.distance {
			t.Error(
				"For", pair.s1, pair.s2,
				"expected", pair.distance,
				"got", v,
			)
		}
	}
}

func TestRun(t *testing.T) {
	r, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	v := run("test", r)

	for i := range testlevenshteinDistance {

		if v[i].Distance != testRun[i].Distance || v[i].Text != testRun[i].Text {
			t.Error(
				"For", v[i].Text, testRun[i].Text,
				"expected", testRun[i].Distance,
				"got", v[i].Distance,
			)
		}
	}
}
