package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Main() {
	strategyReader := NewFileScanner("day2/data.txt")

	totalscore := 0
	totalCorrectScore := 0

	for strategyReader.Scan() {
		hands := strings.Split(strategyReader.Text(), " ")
		theirHand := hands[0]
		myHand := hands[1]
		totalscore += calcHandScore(myHand)
		totalscore += calcResultScore(theirHand, myHand)

		targetResult := myHand
		myHand = calculateNeededHand(theirHand, targetResult)
		totalCorrectScore += calcHandScore(myHand)
		totalCorrectScore += calcResultScore(theirHand, myHand)
	}

	fmt.Printf("The total score of the strategy card is %d", totalscore)
	fmt.Println()
	fmt.Printf("The total score of the strategy card, when read correctly, is %d", totalCorrectScore)
	fmt.Println()

}

func NewFileScanner(filePath string) *bufio.Scanner {
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	return fileScanner //TODO: CLOSE FILE AFTER READING
}

func calcHandScore(hand string) int {
	handPoints := make(map[string]int)
	handPoints["X"] = 1
	handPoints["Y"] = 2
	handPoints["Z"] = 3

	return handPoints[hand]
}

func calcResultScore(theirHand, myHand string) int {
	needToBeat := make(map[string]string)
	equals := make(map[string]string)
	needToBeat["A"] = "Y"
	needToBeat["B"] = "Z"
	needToBeat["C"] = "X"

	equals["A"] = "X"
	equals["B"] = "Y"
	equals["C"] = "Z"

	if myHand == equals[theirHand] {
		return 3
	}

	if myHand == needToBeat[theirHand] {
		return 6
	}

	return 0

}

func calculateNeededHand(theirHand, targetResult string) string {
	needToBeat := make(map[string]string)
	equals := make(map[string]string)
	needToLose := make(map[string]string)

	needToBeat["A"] = "Y"
	needToBeat["B"] = "Z"
	needToBeat["C"] = "X"

	needToLose["A"] = "Z"
	needToLose["B"] = "X"
	needToLose["C"] = "Y"

	equals["A"] = "X"
	equals["B"] = "Y"
	equals["C"] = "Z"

	if targetResult == "Y" {
		return equals[theirHand]
	}

	if targetResult == "Z" {
		return needToBeat[theirHand]
	}

	return needToLose[theirHand]
}
