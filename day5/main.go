package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	lines, err := getInput("sample-input")
	if err != nil {
		log.Fatal(err)
	}

	crates, err := parseCrates(lines[0:3], 3)
	if err != nil {
		log.Fatal(err)
	}

	for _, crate := range crates {
		println(crate)
	}
}

func getInput(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return strings.Split(string(f), "\r\n"), err
}

// source: https://stackoverflow.com/questions/28058278/how-do-i-reverse-a-slice-in-go
func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func parseCrates(lines []string, numCols int) ([]string, error) {
	if len(lines) == 0 {
		return nil, nil
	}

	// start from bottom
	ReverseSlice(lines)
	crates := make([]string, numCols)

	for _, line := range lines {
		col := 0

		for i := 0; i < len(line); i += 3 {
			i++

			if string(line[i]) == " " {
				col++
				continue
			}

			crate := string(line[i])
			crates[col] += crate
			col++
		}
	}

	return crates, nil
}
