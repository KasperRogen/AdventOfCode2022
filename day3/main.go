package day3

import (
	"advent_of_code/shared"
	"bufio"
	"fmt"
	"unicode"
)

func Main() {
	ruckSackScanner := shared.NewFileScanner("day3/data.txt")
	part1(ruckSackScanner)
	ruckSackScanner = shared.NewFileScanner("day3/data.txt")
	part2(ruckSackScanner)
}

func part2(ruckSackScanner *bufio.Scanner) {
	itemSum := 0
	elfIndex := 0
	var badgeCandidates []rune

	for ruckSackScanner.Scan() {
		rucksack := ruckSackScanner.Text()
		elfGroupIndex := elfIndex % 3

		if elfGroupIndex == 0 {
			badgeCandidates = []rune(rucksack)
			elfIndex++
			continue
		}

		var currentBadgeCandidates []rune
		for _, ruckSackItem := range rucksack {
			for _, badgeCandidate := range badgeCandidates {
				if ruckSackItem == badgeCandidate {
					currentBadgeCandidates = append(currentBadgeCandidates, badgeCandidate)
				}
			}
		}

		badgeCandidates = currentBadgeCandidates

		if elfGroupIndex == 2 {
			badge := badgeCandidates[0]
			if unicode.IsUpper(badge) {
				itemSum += int(badge - 38)
			}

			if unicode.IsLower(badge) {
				itemSum += int(badge - 96)
			}
		}
		elfIndex++
	}

	fmt.Printf("The sum of badges for part 2 is %d\n", itemSum)
}

func part1(ruckSackScanner *bufio.Scanner) {
	itemSum := 0
	for ruckSackScanner.Scan() {
		rucksack := ruckSackScanner.Text()
		rucksackSize := len(rucksack)
		compartment1 := rucksack[0 : rucksackSize/2]
		compartment2 := rucksack[rucksackSize/2 : rucksackSize]

		var sharedItem rune

		for _, charToFind := range compartment1 {
			for _, currentChar := range compartment2 {
				if charToFind == currentChar {
					sharedItem = currentChar
					break
				}
			}
		}
		if unicode.IsUpper(sharedItem) {
			itemSum += int(sharedItem - 38)
		}

		if unicode.IsLower(sharedItem) {
			itemSum += int(sharedItem - 96)
		}

	}

	fmt.Printf("The total sum of shared items for part 1 of the puzzle is: %d\n", itemSum)
}
