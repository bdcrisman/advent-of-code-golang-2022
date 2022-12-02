package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	caloriesPerElf := parseCaloriesPerElf(content)
	highestCalorieElf := highestCaloricElf(caloriesPerElf)

	if highestCalorieElf < 0 {
		println("no elf")
		return
	}

	fmt.Printf("Highest caloric elf is elf #%d with %d calories\n", highestCalorieElf, caloriesPerElf[highestCalorieElf])
}

func getInput(path string) (string, error) {
	f, err := ioutil.ReadFile("input")
	if err != nil {
		println(err.Error())
		return "", err
	}

	return string(f), err
}

func parseCaloriesPerElf(content string) []int {
	lines := strings.Split(content, "\r\n")
	if len(lines) == 0 {
		return nil
	}

	caloriesPerElf := make([]int, 0)
	caloricSum := 0
	for _, line := range lines {
		if line == "" {
			caloriesPerElf = append(caloriesPerElf, caloricSum)
			caloricSum = 0
			continue
		}

		v, _ := strconv.Atoi(line)
		caloricSum += v
	}

	return caloriesPerElf
}

func highestCaloricElf(caloriesPerElf []int) int {
	if len(caloriesPerElf) <= 0 {
		return -1
	}

	index := 0
	max := 0
	for i, v := range caloriesPerElf {
		if v > max {
			max = v
			index = i
		}
	}

	return index
}
