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

	println(len(lines))
}

func getInput(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	return strings.Split(string(f), "\r\n"), err
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
