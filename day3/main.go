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
	lines, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	createAlphaMaps()
	Part1(lines)
	Part2(lines)
}

func getInput(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return strings.Split(string(f), "\r\n"), err
}

func Part1(lines []string) {
	sum := 0
	for _, line := range lines {
		sum += parsePriorityPart1(line)
	}

	fmt.Printf("Part 1 priority sum: %d\n", sum)
}

func Part2(lines []string) {
	sum := 0

	for i := 0; i < len(lines); i += 3 {
		sum += parsePriorityPart2(lines[i], lines[i+1], lines[i+2])

	}

	fmt.Printf("Part 2 priority sum: %d\n", sum)
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

func parsePriorityPart1(line string) int {
	halfLen := len(line) / 2
	firstHalf := line[0:halfLen]
	secondHalf := line[halfLen:]

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

func parsePriorityPart2(line1, line2, line3 string) int {
	for _, x := range line1 {
		for _, y := range line2 {
			for _, z := range line3 {
				if x != y || x != z {
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
	}

	return 0
}
