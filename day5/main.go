package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	crates := parseCrates(lines[0:8], 9)
	if len(crates) == 0 {
		log.Fatal("no crates")
	}

	// src: https://stackoverflow.com/questions/27055626/concisely-deep-copy-a-slice
	cratesPart1 := make([]string, len(crates))
	cratesPart2 := make([]string, len(crates))
	copy(cratesPart1, crates)
	copy(cratesPart2, crates)

	moves := parseMoves(lines[10:])
	if len(moves) == 0 {
		log.Fatal("no moves")
	}

	Part1(cratesPart1, moves)
	Part2(cratesPart2, moves)
}

func getInput(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return strings.Split(string(f), "\r\n"), err
}

func Part1(crates []string, moves []int) {
	movedCrates := moveCratesPart1(crates, moves)
	topCrates := getTopChars(movedCrates)
	fmt.Printf("Part 1, top crates: %s\n", topCrates)
}

func Part2(crates []string, moves []int) {
	crates = moveCratesPart2(crates, moves)
	topCrates := getTopChars(crates)
	fmt.Printf("Part 2, top crates: %s\n", topCrates)
}

// source: https://stackoverflow.com/questions/28058278/how-do-i-reverse-a-slice-in-go
func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}

func parseCrates(lines []string, numCols int) []string {
	if len(lines) == 0 {
		return nil
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

	return crates
}

func parseMoves(lines []string) []int {
	if len(lines) == 0 {
		return nil
	}

	moves := make([]int, 0)
	for _, line := range lines {
		if line == " " {
			continue
		}

		splits := strings.Split(line, " ")
		move, _ := strconv.Atoi(splits[1])
		from, _ := strconv.Atoi(splits[3])
		to, _ := strconv.Atoi(splits[5])

		moves = append(moves, move, from-1, to-1)
	}

	return moves
}

func moveCratesPart1(crates []string, moves []int) []string {
	for i := 0; i < len(moves); i += 3 {
		nCrates := moves[i]
		from := moves[i+1]
		to := moves[i+2]

		for j := 0; j < nCrates; j++ {
			crates[to] += string(crates[from][len(crates[from])-1])
			crates[from] = string(crates[from][:len(crates[from])-1])
		}
	}

	return crates
}

func moveCratesPart2(crates []string, moves []int) []string {
	for i := 0; i < len(moves); i += 3 {
		nCrates := moves[i]
		from := moves[i+1]
		to := moves[i+2]

		pile := ""
		for j := 0; j < nCrates; j++ {
			lenFrom := len(crates[from]) - 1
			pile += string(crates[from][lenFrom])
			crates[from] = string(crates[from][:lenFrom])
		}
		pile = reverseString(pile)
		crates[to] += pile
	}

	return crates
}

// src: https://www.geeksforgeeks.org/how-to-reverse-a-string-in-golang/
func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func removeLastRune(s string) string {
	r := []rune(s)
	return string(r[:len(r)-1])
}

func getTopChars(crates []string) string {
	s := ""
	for _, crate := range crates {
		s += string(crate[len(crate)-1])
	}
	return s
}
