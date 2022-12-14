package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

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
	sum := parseAssignmentPairs(lines, true)
	fmt.Printf("Part 1 sum: %d\n", sum)
}

func Part2(lines []string) {
	sum := parseAssignmentPairs(lines, false)
	fmt.Printf("Part 2 sum: %d\n", sum)
}

func parseAssignmentPairs(lines []string, isPart1 bool) int {
	sum := 0
	for _, line := range lines {
		elf1, elf2 := parsePair(line)
		if elf1 == nil || elf2 == nil {
			continue
		}

		if isPart1 && containsSlice(elf1, elf2) {
			sum++
			continue
		}

		if !isPart1 && slicesOverlap(elf1, elf2) {
			sum++
		}
	}

	return sum
}

func parsePair(line string) ([]int, []int) {
	splits := strings.Split(line, ",")
	if len(splits) != 2 {
		return nil, nil
	}

	sequence := strings.Split(splits[0], "-")
	if len(sequence) != 2 {
		return nil, nil
	}

	firstN, err := strconv.Atoi(sequence[0])
	if err != nil {
		return nil, nil
	}
	secondN, err := strconv.Atoi(sequence[1])
	if err != nil {
		return nil, nil
	}

	arr1 := make([]int, 0)
	for i := firstN; i <= secondN; i++ {
		arr1 = append(arr1, i)
	}

	sequence = strings.Split(splits[1], "-")
	firstN, err = strconv.Atoi(sequence[0])
	if err != nil {
		return nil, nil
	}
	secondN, err = strconv.Atoi(sequence[1])
	if err != nil {
		return nil, nil
	}

	arr2 := make([]int, 0)
	for i := firstN; i <= secondN; i++ {
		arr2 = append(arr2, i)
	}

	return arr1, arr2
}

func containsSlice(slice1, slice2 []int) bool {
	if isEmptyOrNil(slice1) || isEmptyOrNil(slice2) {
		return false
	}

	if len(slice1) == len(slice2) && slice1[0] == slice2[0] {
		return true
	} else if len(slice1) > len(slice2) {
		return slice2[0] >= slice1[0] && slice2[len(slice2)-1] <= slice1[len(slice1)-1]
	} else if len(slice2) > len(slice1) {
		return slice1[0] >= slice2[0] && slice1[len(slice1)-1] <= slice2[len(slice2)-1]
	} else {
		return false
	}
}

func slicesOverlap(slice1, slice2 []int) bool {
	if isEmptyOrNil(slice1) || isEmptyOrNil(slice2) {
		return false
	}

	if len(slice1) == len(slice2) && slice1[0] == slice2[0] {
		return true
	}

	if len(slice1) >= len(slice2) {
		return isWithinSlice(slice1, slice2[0]) ||
			isWithinSlice(slice1, slice2[len(slice2)-1])
	}

	if len(slice2) >= len(slice1) {
		return isWithinSlice(slice2, slice1[0]) ||
			isWithinSlice(slice2, slice1[len(slice1)-1])
	}

	return false
}

func isEmptyOrNil(slice []int) bool {
	return len(slice) == 0
}

func isWithinSlice(slice []int, value int) bool {
	return value >= slice[0] && value <= slice[len(slice)-1]
}
