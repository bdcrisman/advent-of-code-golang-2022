package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	input, err := getInput("sample-input")
	if err != nil {
		log.Fatal(err)
	}

	Part1(input, 4)
	Part2(input, 14)
}

func getInput(path string) (string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return "", err
	}
	return string(f), err
}

func Part1(input string, dataLen int) {
	marker := findMarker(input, dataLen)
	fmt.Printf("Part 1 marker: %d\n", marker)
}

func Part2(input string, dataLen int) {
	marker := findMarker(input, dataLen)
	fmt.Printf("Part 2 marker: %d\n", marker)
}

func findMarker(input string, dataLen int) int {
	substr := ""

	for i, r := range input {
		substr += string(r)
		if isMarker(substr, dataLen) {
			return i
		}
	}

	return -1
}

func isMarker(s string, dataLen int) bool {
	if len(s) < dataLen {
		return false
	}

	m := make(map[string]bool)
	count := 0

	for i := len(s) - 1; count <= dataLen; i-- {
		key := string(s[i])
		if _, keyExists := m[key]; keyExists {
			return false
		} else {
			m[key] = true
		}
		count++
	}

	return true
}
