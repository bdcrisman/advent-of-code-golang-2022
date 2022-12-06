package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
)

var (
	LowerAlphas map[string]int
	UpperAlphas map[string]int
)

func main() {
	contentArr, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	createAlphaMaps()
	Part1(contentArr)
}

func getInput(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return strings.Split(string(f), "\r\n"), err
}

func Part1(contentArr []string) {
	sum := 0
	for _, line := range contentArr {
		sum += parseLinePriority(line)
	}

	fmt.Printf("Part 1 priority sum: %d\n", sum)
}

func createAlphaMaps() {
	LowerAlphas = make(map[string]int)
	UpperAlphas = make(map[string]int)

	i := 1
	for r := 'a'; r <= 'z'; r++ {
		R := unicode.ToUpper(r)

		LowerAlphas[string(r)] = i
		UpperAlphas[string(R)] = i + 26

		i++
	}
}

func parseLinePriority(line string) int {
	halfLen := len(line) / 2
	firstHalf := line[0:halfLen]
	secondHalf := line[halfLen:len(line)]

	for _, x := range firstHalf {
		for _, y := range secondHalf {
			if x != y {
				continue
			}

			priority := 0

			if unicode.IsUpper(x) {
				priority = UpperAlphas[string(x)]
			} else {
				priority = LowerAlphas[string(x)]
			}

			return priority
		}
	}

	return 0
}
