package day6

import (
	"advent_of_code/shared"
	"fmt"
)

func Main() {
	streamScanner := shared.NewFileScanner("day6/data.txt")
	streamScanner.Scan()
	line := streamScanner.Text()

	var currentMarker []rune
	currentIndex := 0
	for i, r := range line {
		indexInMarker := indexOf(currentMarker, r)
		if indexInMarker != -1 {
			currentIndex = i - (len(currentMarker) - indexInMarker)
			currentMarker = currentMarker[indexInMarker+1:]
		}

		currentMarker = append(currentMarker, r)

		if len(currentMarker) == 4 {
			break
		}
	}

	markerIndex := currentIndex

	var currentMsgMarker []rune
	currentIndex = 0
	for i, r := range line {
		indexInMarker := indexOf(currentMsgMarker, r)
		if indexInMarker != -1 {
			currentIndex = i - (len(currentMsgMarker) - indexInMarker)
			currentMsgMarker = currentMsgMarker[indexInMarker+1:]
		}

		currentMsgMarker = append(currentMsgMarker, r)

		if len(currentMsgMarker) == 14 {
			break
		}
	}

	messageIndex := currentIndex

	//Elves aren't 0-index (noobs), so we add 5 instead of 4.
	fmt.Println("The index of the stream marker is", markerIndex+5)
	fmt.Println("The index of the message marker is", messageIndex+15)

}

func indexOf(slice []rune, r rune) int {
	sliceLength := len(slice)

	if sliceLength == 0 {
		return -1
	}

	for i := 0; i < sliceLength; i++ {
		if slice[i] == r {
			return i
		}
	}

	return -1
}
