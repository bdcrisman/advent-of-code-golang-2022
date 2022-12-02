package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

/*
	### PART 1
	A|X: Rock		::	1
	B|Y: Paper		::	2
	C|Z: Scissors	::	3

	### PART 2
	X: Loss
	Y: Draw
	Z: Win
*/

const (
	WinPoints  = 6
	DrawPoints = 3
	LossPoints = 0

	Rock     = "A"
	Paper    = "B"
	Scissors = "C"

	Loss = "X"
	Draw = "Y"
	Win  = "Z"
)

func main() {
	contentArr, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	Part1(contentArr)
	Part2(contentArr)
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
	score := 0
	for _, play := range contentArr {
		score += getPlayScore(play, true)
	}

	fmt.Printf("Final Score Part 1: %d\n", score)
}

func Part2(contentArr []string) {
	score := 0
	for _, play := range contentArr {
		score += getPlayScore(play, false)
	}
	fmt.Printf("Final Score Part 2: %d", score)
}

func getPlayScore(play string, isPart1 bool) int {
	plays := strings.Split(play, " ")
	if len(plays) != 2 {
		return 0
	}

	if isPart1 {
		return getWinnerPart1(plays[0], convertPlay(plays[1]))
	}

	return getWinnerPart2(plays[0], plays[1])
}

func getWinnerPart1(opp string, player string) int {
	playerScore := choiceScore(player)

	if opp == player {
		return DrawPoints + playerScore
	} else if opp == Rock && player == Paper ||
		opp == Paper && player == Scissors ||
		opp == Scissors && player == Rock {
		return WinPoints + playerScore
	} else if opp == Rock && player == Scissors ||
		opp == Paper && player == Rock ||
		opp == Scissors && player == Paper {
		return LossPoints + playerScore
	} else {
		return 0
	}
}

func getWinnerPart2(opp string, player string) int {
	if player == Loss {
		return LossPoints + loseToOpp(opp)
	} else if player == Draw {
		return DrawPoints + choiceScore(opp)
	} else {
		return WinPoints + defeatOpp(opp)
	}
}

func defeatOpp(opp string) int {
	score := 0

	switch opp {
	case Rock:
		score = choiceScore(Paper)
	case Paper:
		score = choiceScore(Scissors)
	case Scissors:
		score = choiceScore(Rock)
	default:
		score = 0
	}

	return score
}

func loseToOpp(opp string) int {
	score := 0

	switch opp {
	case Rock:
		score = choiceScore(Scissors)
	case Paper:
		score = choiceScore(Rock)
	case Scissors:
		score = choiceScore(Paper)
	default:
		score = 0
	}

	return score
}

func convertPlay(player string) string {
	if player == "X" {
		return Rock
	} else if player == "Y" {
		return Paper
	} else {
		return Scissors
	}
}

func choiceScore(s string) int {
	result := 0

	switch s {
	case Rock, "X":
		result = 1

	case Paper, "Y":
		result = 2

	case Scissors, "Z":
		result = 3

	default:
		result = 0
	}

	return result
}
