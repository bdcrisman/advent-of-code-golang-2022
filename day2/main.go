package main

import (
	"io/ioutil"
	"log"
	"strings"
)

/*
	A|X: Rock		::	1
	B|Y: Paper		::	2
	C|Z: Scissors	::	3
*/

const Win = 6
const Draw = 3
const Loss = 0

type RockPaperScissorsResult struct {
	Opponent       string
	PlayerResponse string
}

func main() {
	contentArr, err := getInput("input")
	if err != nil {
		log.Fatal(err)
	}

	finalScore := playGame(contentArr)
	println(finalScore)
}

func getInput(path string) ([]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		println(err.Error())
		return nil, err
	}

	return strings.Split(string(f), "\r\n"), err
}

func playGame(contentArr []string) int {
	sum := 0
	for _, play := range contentArr {
		score := getPlayScore(play)
		if score < 0 {
			continue
		}

		sum += score
	}

	return sum
}

func getPlayScore(play string) int {
	plays := strings.Split(play, " ")
	if len(plays) != 2 {
		return -1
	}

	return getScore(plays[0], plays[1])
}

func getScore(opp string, player string) int {
	return getWinner(opp, convertPlay(player))
}

func getWinner(opp string, player string) int {
	playerScore := choiceScore(player)

	if opp == player {
		return Draw + playerScore
	} else if opp == "A" && player == "B" ||
		opp == "B" && player == "C" ||
		opp == "C" && player == "A" {
		return Win + playerScore
	} else if opp == "A" && player == "C" ||
		opp == "B" && player == "A" ||
		opp == "C" && player == "B" {
		return Loss + playerScore
	} else {
		return 0
	}
}

func convertPlay(player string) string {
	if player == "X" {
		return "A"
	} else if player == "Y" {
		return "B"
	} else {
		return "C"
	}
}

func choiceScore(s string) int {
	result := 0

	switch s {
	case "A", "X":
		result = 1

	case "B", "Y":
		result = 2

	case "C", "Z":
		result = 3

	default:
		result = 0
	}

	return result
}
