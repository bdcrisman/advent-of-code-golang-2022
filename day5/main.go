package main

import (
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

	moves := parseMoves(lines[10:])
	if len(moves) == 0 {
		log.Fatal("no moves")
	}

	crates = moveCrates(crates, moves)
	topCrates := getTopChars(crates)

	println(topCrates)
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

func moveCrates(crates []string, moves []int) []string {
	for i := 0; i < len(moves); i += 3 {
		nCrates := moves[i]
		from := moves[i+1]
		to := moves[i+2]

		// fmt.Printf("%d: %d: %d\n", nCrates, from, to)
		for j := 0; j < nCrates; j++ {
			crates[to] += string(crates[from][len(crates[from])-1])
			crates[from] = string(crates[from][:len(crates[from])-1]) //removeLastRune(crates[from])

			// fmt.Printf("to: %s :: from: %s\n", crates[to], crates[from])
		}
	}

	return crates
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
