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
	content, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	caloriesPerElf := parseCaloriesPerElf(content)
	Day1(caloriesPerElf)
	Day2(caloriesPerElf, 3)
}

func Day1(caloriesPerElf []int) {
	highestCalorieElf := highestCaloricElf(caloriesPerElf)

	if highestCalorieElf < 0 {
		println("no elf")
		return
	}

	fmt.Printf("Highest caloric elf is elf #%d with %d calories\n", highestCalorieElf, caloriesPerElf[highestCalorieElf])
}

func Day2(caloriesPerElf []int, nElves int) {
	sumOfTopElves := sumTopElvesCalories(caloriesPerElf, nElves)
	fmt.Printf("Top %d elves have a total of %d calories", nElves, sumOfTopElves)
}

func getInput(path string) (string, error) {
	f, err := ioutil.ReadFile(path)
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

func sumTopElvesCalories(caloriesPerElf []int, nElves int) int {
	if len(caloriesPerElf) <= 0 || nElves > len(caloriesPerElf) {
		return -1
	}

	sort.Slice(caloriesPerElf, func(i, j int) bool {
		return caloriesPerElf[i] > caloriesPerElf[j]
	})

	sum := 0
	for i := 0; i < nElves; i++ {
		sum += caloriesPerElf[i]
	}

	return sum
}
