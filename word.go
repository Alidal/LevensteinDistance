package main

type Word struct {
	Text     string
	Distance int
}

type Words []Word

func (slice Words) Len() int {
	return len(slice)
}

func (slice Words) Less(i, j int) bool {
	return slice[i].Distance < slice[j].Distance
}

func (slice Words) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
